//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package utils

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/uber/aresdb/common"
	"os"
)

// bindEnvironments binds environment variables to viper
func bindEnvironments(v *viper.Viper) {
	v.SetEnvPrefix("ares")
	v.BindEnv("env")
}

// AddFlags adds flags to command
func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String("config", "config/ares.yaml", "Ares config file")
	cmd.Flags().IntP("port", "p", 0, "Ares service port")
	cmd.Flags().IntP("debug_port", "d", 0, "Ares service debug port")
	cmd.Flags().StringP("root_path", "r", "ares-root", "Root path of the data directory")
	cmd.Flags().Bool("scheduler_off", false, "Start server with scheduler off, no archiving and backfill")
}

// ReadConfig populate AresServerConfig
func ReadConfig(defaultCfg map[string]interface{}, flags *pflag.FlagSet) (common.AresServerConfig, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	// bind command flags
	v.BindPFlags(flags)

	bindEnvironments(v)

	// set defaults
	v.SetDefault("root_path", "ares-root")
	hostname, err := os.Hostname()
	if err != nil {
		panic(StackError(err, "cannot get host name"))
	}
	v.SetDefault("cluster", map[string]interface{}{
		"instance_name": hostname,
	})
	v.MergeConfigMap(defaultCfg)

	// merge in config file
	if cfgFile, err := flags.GetString("config"); err == nil && cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("ares")
		v.AddConfigPath("./config")
	}

	if err := v.MergeInConfig(); err == nil {
		fmt.Println("Using config file: ", v.ConfigFileUsed())
	}

	var cfg common.AresServerConfig
	err = v.Unmarshal(&cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	})
	return cfg, err
}

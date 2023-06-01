package cmd

import (
	
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"zgwldrc/rmqtopsync/mod"
)

var cfgFile string
var rootCmdOptSrc string
var rootCmdOptDst string
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rmqtopsync",
	Short: "sync topics from src rocketmq console url to another dest console url",
	Long: `sync topics from src rocketmq console url to another dest console url`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		src := mod.NewRocketMQConsole(rootCmdOptSrc)
		dst := mod.NewRocketMQConsole(rootCmdOptDst)
		topics := src.GetTopics()
		fmt.Printf("%s\n", topics)
		resp := src.GetClusters()
		clusterName := ""
		brokerNames := []string{}
		for k, v := range resp.Data.ClusterInfo.ClusterAddrTable {
			clusterName = k
			brokerNames = v
			break
		}
		
		for _, t := range topics {
			req := &mod.CreateOrUpdateReq{
				WriteQueueNums: 16,
				ReadQueueNums: 16,
				Perm: 6,
				Order: false,
				TopicName: t,
				BrokerNameList: brokerNames,
				ClusterNameList: []string{clusterName},
			}
			resp := dst.CreateOrUpdate(req)
			if resp.Status == 0 {
				fmt.Printf("topic %s create or update success!", t)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rmqtopsync.yaml)")
	rootCmd.Flags().StringVarP(&rootCmdOptSrc, "src", "s", "http://localhost:8080", "src rocketmq console url, example http://rmqconsole.example.com")
	rootCmd.Flags().StringVarP(&rootCmdOptDst, "dst", "d", "http://localhost:8080", "dst rocketmq console url, example http://rmqconsole.example.com")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rmqtopsync" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rmqtopsync")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

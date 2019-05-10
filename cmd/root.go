package cmd

import (
  "github.com/spf13/cobra"
  "github.com/reactiveops/dd-manager/conf"
  "github.com/reactiveops/dd-manager/pkg/controller"
  "github.com/reactiveops/dd-manager/pkg/util"
  log "github.com/sirupsen/logrus"
  "os"
)


func RootCmd() *cobra.Command {
  root := &cobra.Command {
    Use: "dd-manager",
    Short: "Kubernetes datadog monitor manager",
    Long: "A kubernetes agent that manages datadog monitors.",
    Run: run,
  }
  return root
}


func loadConfig(cmd *cobra.Command)(*conf.Config) {
  log.SetReportCaller(true)
  log.SetOutput(os.Stdout)

  config := conf.New()
  return config
}


func run(cmd *cobra.Command, args []string) {
  conf := loadConfig(cmd)
  controller.NewController(conf, util.GetKubeClient())
}

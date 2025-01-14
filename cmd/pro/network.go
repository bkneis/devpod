package pro

import (
	"context"
	"os"

	"github.com/loft-sh/devpod/cmd/pro/flags"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
	"tailscale.com/tsnet"
)

// SelfCmd holds the cmd flags
type NetworkCmd struct {
	*flags.GlobalFlags
	Log log.Logger

	Host string
}

// NewSelfCmd creates a new command
func NewNetworkCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &NetworkCmd{
		GlobalFlags: globalFlags,
		Log:         log.GetInstance(),
	}
	c := &cobra.Command{
		Use:    "network",
		Short:  "Setup tsnet server on the workspace to expose SSH port to the tail net",
		Hidden: true,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			return cmd.Run(cobraCmd.Context())
		},
	}

	c.Flags().StringVar(&cmd.Host, "host", "", "The pro instance to use")
	_ = c.MarkFlagRequired("host")

	return c
}

func (cmd *NetworkCmd) Run(ctx context.Context) error {
	// devPodConfig, err := config.LoadConfig(cmd.Context, cmd.Provider)
	// if err != nil {
	// 	return err
	// }

	// _, err = platform.ProviderFromHost(ctx, devPodConfig, cmd.Host, cmd.Log)
	// if err != nil {
	// 	return fmt.Errorf("load provider: %w", err)
	// }

	s := &tsnet.Server{
		Hostname:     cmd.Host,
		RunWebClient: true,
		ControlURL:   "http://loft.devpod-pro.svc.cluster.local/coordinator/",
		AuthKey:      os.Getenv("TS_AUTHKEY"),
	}

	defer s.Close()

	err := s.Start()
	if err != nil {
		cmd.Log.Fatal(err)
	}

	for {
	}
}

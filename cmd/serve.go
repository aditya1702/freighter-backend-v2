package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stellar/freighter-backend-v2/internal/config"
)

type serveCmd struct {
	cfg *config.Config
}

func (s *serveCmd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "serve",
		Short:         "Start the server",
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return s.Run()
		},
	}

	// RPC Config
	cmd.Flags().StringVar(&s.cfg.RpcConfig.RpcPubnetURL, "rpc-pubnet-url", "", "The URL of the pubnet RPC")
	cmd.Flags().StringVar(&s.cfg.RpcConfig.RpcTestnetURL, "rpc-testnet-url", "https://horizon-testnet.stellar.org", "The URL of the testnet RPC")

	// Horizon Config
	cmd.Flags().StringVar(&s.cfg.HorizonConfig.HorizonPubnetURL, "horizon-pubnet-url", "https://horizon.stellar.org", "The URL of the pubnet Horizon")
	cmd.Flags().StringVar(&s.cfg.HorizonConfig.HorizonTestnetURL, "horizon-testnet-url", "https://horizon-testnet.stellar.org", "The URL of the testnet Horizon")

	// Redis Config
	cmd.Flags().StringVar(&s.cfg.RedisConfig.ConnectionName, "redis-connection-name", "freighter-redis", "The name of the Redis connection")
	cmd.Flags().StringVar(&s.cfg.RedisConfig.Host, "redis-host", "localhost", "The Redis host")
	cmd.Flags().IntVar(&s.cfg.RedisConfig.Port, "redis-port", 6379, "The Redis port")

	// Prices Config
	cmd.Flags().IntVar(&s.cfg.PricesConfig.BatchUpdateDelayMilliseconds, "prices-batch-update-delay", 5000, "Delay between batch updates in milliseconds")
	cmd.Flags().IntVar(&s.cfg.PricesConfig.CalculationTimeoutMilliseconds, "prices-calculation-timeout", 10000, "Timeout for price calculations in milliseconds")
	cmd.Flags().IntVar(&s.cfg.PricesConfig.UpdateIntervalMilliseconds, "prices-update-interval", 30000, "Interval between price updates in milliseconds")
	cmd.Flags().IntVar(&s.cfg.PricesConfig.UpdateBatchSize, "prices-update-batch-size", 50, "Size of price update batches")
	cmd.Flags().IntVar(&s.cfg.PricesConfig.PriceStalenessThreshold, "prices-staleness-threshold", 300000, "Threshold for price staleness")

	// Blockaid Config
	cmd.Flags().BoolVar(&s.cfg.BlockaidConfig.UseBlockaidDappScanning, "use-blockaid-dapp-scanning", false, "Enable Blockaid dapp scanning")
	cmd.Flags().BoolVar(&s.cfg.BlockaidConfig.UseBlockaidTxScanning, "use-blockaid-tx-scanning", false, "Enable Blockaid transaction scanning")
	cmd.Flags().BoolVar(&s.cfg.BlockaidConfig.UseBlockaidAssetScanning, "use-blockaid-asset-scanning", false, "Enable Blockaid asset scanning")
	cmd.Flags().BoolVar(&s.cfg.BlockaidConfig.UseBlockaidAssetWarningReporting, "use-blockaid-asset-warning-reporting", false, "Enable Blockaid asset warning reporting")
	cmd.Flags().BoolVar(&s.cfg.BlockaidConfig.UseBlockaidTransactionWarningReporting, "use-blockaid-transaction-warning-reporting", false, "Enable Blockaid transaction warning reporting")

	// Coinbase Config
	cmd.Flags().StringVar(&s.cfg.CoinbaseConfig.CoinbaseAPIKey, "coinbase-api-key", "", "Coinbase API key")
	cmd.Flags().StringVar(&s.cfg.CoinbaseConfig.CoinbaseAPISecret, "coinbase-api-secret", "", "Coinbase API secret")
	return cmd
}

func (s *serveCmd) Run() error {
	// Print out config values to verify flag parsing
	fmt.Printf("RPC Pubnet URL: %s\n", s.cfg.RpcConfig.RpcPubnetURL)
	fmt.Printf("Horizon Testnet URL: %s\n", s.cfg.HorizonConfig.HorizonTestnetURL)
	fmt.Printf("Redis Host: %s\n", s.cfg.RedisConfig.Host)
	fmt.Printf("Use Blockaid Dapp Scanning: %v\n", s.cfg.BlockaidConfig.UseBlockaidDappScanning)

	return nil
}

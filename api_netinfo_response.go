package main

import "time"

type ApiNetinfoResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Listening bool     `json:"listening"`
		Listeners []string `json:"listeners"`
		NPeers    string   `json:"n_peers"`
		Peers     []struct {
			NodeInfo struct {
				ProtocolVersion struct {
					P2P   string `json:"p2p"`
					Block string `json:"block"`
					App   string `json:"app"`
				} `json:"protocol_version"`
				Id         string `json:"id"`
				ListenAddr string `json:"listen_addr"`
				Network    string `json:"network"`
				Version    string `json:"version"`
				Channels   string `json:"channels"`
				Moniker    string `json:"moniker"`
				Other      struct {
					TxIndex    string `json:"tx_index"`
					RpcAddress string `json:"rpc_address"`
				} `json:"other"`
			} `json:"node_info"`
			IsOutbound       bool `json:"is_outbound"`
			ConnectionStatus struct {
				Duration    string `json:"Duration"`
				SendMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"SendMonitor"`
				RecvMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"RecvMonitor"`
				Channels []struct {
					ID                int    `json:"ID"`
					SendQueueCapacity string `json:"SendQueueCapacity"`
					SendQueueSize     string `json:"SendQueueSize"`
					Priority          string `json:"Priority"`
					RecentlySent      string `json:"RecentlySent"`
				} `json:"Channels"`
			} `json:"connection_status"`
			RemoteIp string `json:"remote_ip"`
		} `json:"peers"`
	} `json:"result"`
}

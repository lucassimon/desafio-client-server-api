/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lucassimon/desafio-client-server-api/internal/dto"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Fetch the dollar price in real",
	Long:  `Consumes the endpoint /cotacao/ and get the dollar price in real converted.`,
	Run: func(cmd *cobra.Command, args []string) {
		// terá um timeout máximo de 300ms para receber o resultado do server.go
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://0.0.0.0:8080/cotacao/", nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var cotacao_resp dto.CreateDollarInput
		err = json.Unmarshal(body, &cotacao_resp)
		if err != nil {
			log.Fatal(err)
		}

		var c dto.CreateDollarOutput
		c.Bid = cotacao_resp.USDBRL.Bid
		c.CreateDate = cotacao_resp.USDBRL.CreateDate
		// terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
		f, err := os.Create("cotacao.txt")
		if err != nil {
			log.Println(err)
		}
		tamanho, err := fmt.Fprintln(f, "Dólar", c.Bid, "at", c.CreateDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

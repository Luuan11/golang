package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "https://www.mercadolivre.com.br/",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs e nomes de Servidor na internet",
			Flags: flags,
			Action: BuscaIPS,
		},
		{
			Name: "servidores",
			Usage: "Busc o nome do servidores na internet",
			Flags: flags,
			Action: BuscarServidores,
		},
	}

	return app
}

func BuscaIPS(c *cli.Context){
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}
}

func BuscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

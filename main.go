package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/kataras/iris/v12"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go runApi(&wg)
	wg.Wait()
}

type Config1 struct {
	Mission string
	Map     string
}

func runApi(wg *sync.WaitGroup) {
	var conf Config1
	tomlSetting, err := Asset("data/settings.toml")
	if err != nil {
		fmt.Print(err)
	}
	if _, err := toml.Decode(string(tomlSetting), &conf); err != nil {
		fmt.Println("can't load file settings", err)
	}

	auto, err := Asset("data/gui")
	if err != nil {
		fmt.Print(err)
	}

	app := iris.Default()

	app.Get("/iauto", func(ctx iris.Context) {
		fmt.Println(ctx)
		cmd := ctx.FormValue("cmd")
		if cmd == "gui" {
			ctx.Header("Content-Type", "application/octet-stream")
			ctx.Header("Content-Disposition", "attachment; filename=i")
			ctx.Write(auto)
		} else if cmd == "login" {
			fmt.Println(ctx)
			ctx.Writef("success")
		} else if cmd == "userinfo" {
			fmt.Println(ctx)
			ctx.Writef("1|0|99999999|999|0%%|0|1")
		} else {
			ctx.Writef("")
		}

	})

	app.Get("/iauto/nier.php", func(ctx iris.Context) {
		ctx.Writef(`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`)
		// ctx.Writef(`LXP5oXA/XIpOLafg/bPUQbUggPMRKC+HggV4i60ngQG/ICPVek045cgNv8eQZCjImi+H0S/luWU3NVOgrNE/g8jIuYc0TuJAyH6Od/W5oqgjCFWWmFVP7B0piJ4XjHsslReAjk2+ybkKRQIkSzo4Vpx2za1UK0JmbrAtaOl0UUkHCFfWbaWBSSW65muDMPQcsEAgbTedjBlKU8nEWsd0y/DVfQXT+R87n53u6taR2b8cIjpl+MjzuBZYRA9Qb/aGulA4w3WZ937jaHEFvjEWvE4pYLQGHLXtOpHpYoMQMxpSMr3MXWlEtJ+19QbGJOf61xDX4cc1nR/n6iLpGJfckmh2ELcNY9CQkfzALEW7RjF34QWi1AXe8BmXqzfVn9nqUX9K6oEAgheqNiVJHGq0xCyJY1yV73CydM0ZuKhWd2AdWfYASYDoaukFdHAsMr0qdLDKn6cxSsFjBds95ZSUpXIIm/VxzVpzVrmEv117TMqw3v6U5xWrORdw12Hqg7aaMeoRKGfUbCUeTn/zDPMD/Nhq8k0cnhqJFt6wjjctAHDUbKsps2BFJzl2LU5TkQ9iWWQOn+MpBQ5QyXR2fC6jtJFbi+Wup6fjYsVQ/ikrXR++oCSwHJNl04bVmQcwQ4nnNc12NKT/dsoHzShS1WaFhFd64w+h3l+EIKtTTuc/noZQdJg44loZ1xBgVVVVxKZVVfNspCfuvMCIPp5Wt50laF8y8Uw4mhNV4Wiz3bZrdTV3+Yz8W8g4SxwCUZ0B5t5dD1/pFS8/VH9YFsurd5LGtSN6R0eUpivrXQPTItmlU++W8aK17yyjjqYb1HjqJcZGqzY4rOmqUl0esavwy4t6gvNMsrQo3rjTGh8sjdune5+MML1PPQhqiZmMTon+d722UmCI+Qg0RNG0/N+pdn3jeFnUyHXTOgn/x9UnBbYWB/CZkcsVyh7WqqMOM0GwDYAjCKYVzwOPdZhjlPv4imEmEOgyX+QcKf3t0vCdaZXscDw8UE3UXAJIWQ42KKhCwKhegYZGBBBDchs84ABrPtrGrZoHQYzrfdUbAaHY5R5LIng1/ZftqQgM+Jiq3JyVOF/23VDYTxKDCVrh6oYRkuyrqdftt2RqdkwcKHgktczyWrJkbPu2JvMwzLnK5e+MJFisqlQRUZjvK6vXZInW8iJ+RUjAJK+fnpQMjtvGe7OFDXeaPnbn/w13MALg3vNm9faPUk3Px8Nlk7p3G/SWD/16y6iNki/9xYMe0vbqIkS1U99nfDRXxuN9EZzdK+9M3Xvk/Rboe2KJW6z93+HVx39TB/IWHFKL/pkqVd321tQwifaJEVIyB1U3yF3VMW8JZBPYaOpB7P+uQu1vocsfy6xqxcgfreAWX7qKSIrIi8TauGo0CwS85CkvXFYC99Th50qjKZkdFur168rj0pwjcA07BXC0i1xo7igx0CcqYjgicZp0trIZKiDABBbON87ZzJ0XOA9E6rB+TsADS+a5qqDShCFR+Ans5D6XmK7fzA+iKhP/s/14paVJuovv4odXPgTzSL4h3RjEv9TRVwgkueabdLu25y82Cb4dYDSthCH+Idboo+6Eb9jKyKRGPzNjM2T71oxwnyyY9JJYHAjkRqZtRutNfFltvmkEKyhf1LjV50rjdpMuLf89eqYk+VDuLd99KvUZZi8ZnYSsCnKWXvhjZV8lWKbhT/YCAtusS6fsI5suoxtutEPFuEDbHLstlxkD0tBN31rPEZbiKkYjCnmhY6FL07ncqVAGHJ288yxsYXafCPDCEA6omOqv/NENWwBFEu5YJIR4DjAw7F2tQALfci2/We4C8Eatdjqwz2bvhvkrtGEhWK7PZ+1iUnbzvsgv1PQ7JczmA5RNoJ/gC5tf/Qg9rHI2SSxqvJoGkns7YKfiNnJvu9ZP1o+YzuJietFRmpfJpswObPdB8AbAXP+JUcp8tU42V55l5xxt5C3TWUvsybDfjgqvuJy65zXSYAYtjzHtjfbauPHzzZKMBf7QkKEoM1SDeLOT7fTQAm0+LhZq2A/p9Ke+3PBDgvEdPDch/C/iLCqbFfCe8kgPv5PVIz2/EQwQ++t7VIy7cwQhN9KK2uYHXqvVRPAzEzRwya8WXhppee0ehhfvj1PIJ0C56is9ZNrb/9rVt98gPxKse2tBPTvVAtDWjJy4A+CaHXqc/rjyq3ONLTZRm5AvAJ/aPmmsW/WOb1vHbGw5j6ZOoOcCZewZ+6EyiqclWgcoMIxFse5o+ZRgyrfI1akenkjIYnIPEyX8pe+fPYdeWstNniOGP3sUlPJLsiELIF3K09/DrjGZR3LsU272XAjMQ4NvsUl8O+Alc5bAW9VHvmymiU6ppGgFINd+l9CQNTfW1yOj24z88n/1+Iag7YDKnA3zPUJavdnHxsS1rQGjq0BWFJBpeKsZ/qjdi1qQJuyDy4e9g+slKTF7/lnl4DNjHfi9QgmEs4xz2LjhYntXhe4UszPTIvW9M5eEGgt93nfP7yIvoP7A2ZiKELj/LaoIy7ms5cyh6EzQ2Zv9wrMfiLdm+VN4LZb1u0GONkudfFPxh/pgA1s+uAgVkB771jONUK63TdmPV14teCVNo30Jr9vkm2D5L1FzxCyLP2U2CA5EXO7aD7FvMjbpQMvtV3xDrc1OQjeI4lZUHbjdyyoaB2H6dnlMZZ7iYfYObUAlnQc45Gjqk1hBPpi2h4nGiUC2pIhQA98sgNrT+JAQkP+Und/2h9wfBHriLGmB28b/w/FAZi07CIyDm1EGQC7kMWXWGE5ptFthNB0/p29I+P/jgUZGLetBsjmg8k4NgHbfaH9SCkfii1B/kDlH4nCIN0F2SqeZeFB7R3VIza7djNR1pxMFzzeSAhOltbgZq82SjxB2YxXJ5Y9SPSgdIW1Z45tg6QgjAmdIBD0sfMMdFQfm8i8E7rPBuV/YoBFQ6cEXB/y0Bc+d6G0wgL4ynr8RlwNJsP38lD0ORviM8wymIzbaa5jj8JzyNMzfyKAgYkrlFbV5ErXhZyGRlhzJEy9PkVsPxrp4AoBlkNkMoALyDL+d2d/nhazrlyrr4vspntAJsO3LF1gA7mTPgf17ivGFs9r+WVEDvT1/iq17phU//XziP0Lsf52OHKojkFbeKS97ZZLyT66l7A2JJDaXBP7BbC0zNczRIiNrzfpkM3/5Tqm4yuRGjxdqXo2adX8KYur9yVqZRO2R5i82aBqT8LDHCytTLxlAoCQyqdZ4ywGDxxiSnAo7DI8HlQBmQygZlbyZu2Y35K2hAZqFsWCJzj/3eO4K6ky9suq0NTTdDD7/uOg3Qwj8WJLtc85Lmg3MTHzTk6p5p4u8z+8nmgZ4tSBp4BAYFCyZBi4Ho0WrnNyITxNXd6fc1qByperDbS3P1OTdjBqpMGekQg8dfwPD4U7eJJCmbtO4LuTXYkQ/WRlghSAzvz32gC69HP9m4N+Pak+RuMWZwxjV1KOGAbUzZ1xjdBGNVAzYh0eB7vK7TeEnAHyMy3Q6dfUL6qTGaBY32o7COdV95CbBXXBBpUPvI85ORwpmX3fkofFvCsxfmEcB0jsq6RON7AnsB0aAmuA71R6+spvISAinvxlis4AXadyyt3bDb/jg1rMocBOCHuUeInqaUPy0nmgnH9BrQUU0GPpqrA8YHGZOWoxCsWUVoZqQTJKKsRZMA/G+RRIZrMSb62+uQWQLpp6jsdpPffhAK3rN+jxylBF6b+EOdesXlhZtnjtepcAQwiIOIEreXYhabouS/72oKuFtptOxq8B9NkwLr+Ue/lvz0dpIHnJYJxs/Io09uYhGMQIQDjSjicRYwyW7PGPCs7J6wsj3vZ6fE964k4G9PJ+sW6d/fri98sYnwKURtmvDq9tuQeYmV9qgNrSoSfnKAFrIbZoj0iMObFZa/5pBFMNcsCoLWf3ZuChXk8P82Ud+WcZtw294d9+2S9mRrWDztpWbASgxevN6tvAyjPkvqL/VwPnufCH/VDkWaaWH8tMIC+yhVOHAe3SW/K4CVJFo6kQ=`)
	})

	app.Get("/iauto/scripts.xml", func(ctx iris.Context) {
		ctx.Writef(conf.Mission)
	})

	app.Get("/iauto/map.ini", func(ctx iris.Context) {
		ctx.Writef(conf.Map)
	})

	certData, err := Asset("data/service.crt")
	if err != nil {
		fmt.Println("Failed to read certificate data", err)
	}

	keyData, err := Asset("data/service.key")
	if err != nil {
		fmt.Println("Failed to read private key data", err)
	}

	go app.Listen(":80")

	app.Run(
		iris.TLS(":443", "", "", func(su *iris.Supervisor) {
			su.Server.TLSConfig = Config(certData, keyData)
		}),
	)

	wg.Done()
	os.Exit(2)
}

func loadCertificate(certificate []byte, privateKey []byte) (tls.Certificate, error) {
	if cert, err := tls.X509KeyPair([]byte(certificate), []byte(privateKey)); err != nil {
		return tls.Certificate{}, err
	} else {
		return cert, nil
	}
}

func Config(certificate []byte, privateKey []byte) *tls.Config {
	certs := make([]tls.Certificate, 1)
	cert, err := loadCertificate(certificate, privateKey)

	if err != nil {
		panic(err)
	}

	certs[0] = cert

	return &tls.Config{
		Certificates: certs,
		NextProtos:   []string{"h2", "http/1.1"},
	}
}

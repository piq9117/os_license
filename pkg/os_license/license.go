package os_license

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Get(licenseType string, fileOutput string) {
	fromAvailableLicense := searchFromAvailableLicense(licenseType)

	res, err := http.Get(githubLicenseApi(fromAvailableLicense))

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalln("Something went wrong with github license api")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var license License
	err = json.Unmarshal(body, &license)
	if err != nil {
		log.Fatalln(err)
	}

	// if file output is provided, print to file.
	// else print to stdout
	if fileOutput != "" {
		f, err := os.Create(fileOutput)

		if err != nil {
			log.Fatalln(err)
		}

		f.WriteString(license.Body)

		defer f.Close()
	} else {
		fmt.Println(license.Body)
	}

}

type License struct {
	Body string `json:"body"`
}

func searchFromAvailableLicense(licenseType string) string {
	var availableLicense [13]string = [13]string{
		"agpl-3.0",
		"apache-2.0",
		"bsd-2-clause",
		"bsd-3-clause",
		"bsl-1.0",
		"cc0-1.0",
		"epl-2.0",
		"gpl-2.0",
		"gpl-3.0",
		"lgpl-2.1",
		"mit",
		"mpl-2.0",
		"unlicense",
	}
	var fromAvailableLicense string
	for _, license := range availableLicense {
		if licenseType == license {
			fromAvailableLicense = license
		}
	}

	if fromAvailableLicense == "" {
		log.Fatalln("License not available")
	}

	return fromAvailableLicense
}

func githubLicenseApi(license string) string {
	var githubLicenseApi = "https://api.github.com/licenses"

	return githubLicenseApi + "/" + license
}

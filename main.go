package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Data struct {
		Catalog struct {
			Category struct {
				ProductsWithMetadata struct {
					List []struct {
						URLPart string `json:"urlPart"`
					} `json:"list"`
				} `json:"productsWithMetadata"`
			} `json:"category"`
		} `json:"catalog"`
	} `json:"data"`
}

type ProductResult struct {
	Data struct {
		Catalog struct {
			Product struct {
				ID      string `json:"id"`
				Title   string `json:"name"`
				URLPart string `json:"urlPart"`
				Media   []struct {
					FullUrl string `json:"fullUrl"`
				} `json:"media"`
				Options []struct {
					Title      string `json:"title"`
					Selections []struct {
						Value string `json:"value"`
					} `json:"selections"`
				} `json:"options"`
			} `json:"product"`
		} `json:"catalog"`
	} `json:"data"`
}

type Product struct {
	ID         string              `json:"id"`
	Title      string              `json:"title"`
	Subtitle   string              `json:"subtitle"`
	ContentURL string              `json:"contentURL"`
	Media      []string            `json:"media"`
	Attributes []map[string]string `json:"attributes"`
}

func main() {
	var aggregatedJSON []Product
	authorizationToken := "brUTfgwc9eaqQ4m_KjbIkjnR-MRt9rGfCLGikGEPiRU.eyJpbnN0YW5jZUlkIjoiMWI0OTQ1ODItZDg5Zi00MmY2LTg0YzAtNTAxOGE3NzI1Y2MyIiwiYXBwRGVmSWQiOiIxMzgwYjcwMy1jZTgxLWZmMDUtZjExNS0zOTU3MWQ5NGRmY2QiLCJtZXRhU2l0ZUlkIjoiN2RlM2ExNjgtNDEyNC00NDljLTg4ZDYtZmViNjkzYWY3NzRjIiwic2lnbkRhdGUiOiIyMDIwLTA5LTIzVDEyOjI3OjE4LjUyOVoiLCJ2ZW5kb3JQcm9kdWN0SWQiOiJQcmVtaXVtMSIsImRlbW9Nb2RlIjpmYWxzZSwiYWlkIjoiOWE0ZjJjNDAtMTIzNC00ZGM3LTg3OWEtMjIzZDMxMzI0N2E1IiwiYmlUb2tlbiI6IjY2YWFlNGVhLTk5YmItMDY2YS0wYzE2LWFlYWUzNGRkMmI4ZSIsInNpdGVPd25lcklkIjoiZmI0Y2Y2ODQtODZkZS00N2E0LWE2NjUtZjE4ZDcxYzA3YzUxIn0"

	URLPartArr, err := getURLPart()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Getting data from api...")

	for _, element := range URLPartArr {
		data, _ := fetchData(element, authorizationToken)
		product := buildProduct(data)
		aggregatedJSON = append(aggregatedJSON, product)
	}

	log.Println("Writing to json file...")
	file, err := json.MarshalIndent(aggregatedJSON, "", " ")

	if err != nil {
		log.Fatalln("Couldn't marshal to file", err)
	}

	if err := ioutil.WriteFile("enhanced.json", file, 0644); err != nil {
		log.Fatalln("Error writing to file", err)
	}
}

func getURLPart() ([]string, error) {
	jsonFile, err := os.Open("lafiancee.json")
	if err != nil {
		log.Println("Could't open the json file")
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Result
	json.Unmarshal([]byte(byteValue), &result)

	productList := result.Data.Catalog.Category.ProductsWithMetadata.List

	var URLPartArr []string
	for _, element := range productList {
		URLPartArr = append(URLPartArr, element.URLPart)
	}

	return URLPartArr, nil
}

func fetchData(productID string, authorizationToken string) ([]byte, error) {
	jsonQuery := map[string]string{
		"query": "query getProductBySlug($externalId: String!, $slug: String!, $withPricePerUnit: Boolean!, $withCountryCodes: Boolean!) {\n          appSettings(externalId: $externalId) {\n            widgetSettings\n      }\n      catalog {\n            product(slug: $slug, onlyVisible: true) {\n                id\n                description\n                isVisible\n                sku\n                ribbon\n                price\n                comparePrice\n                discountedPrice\n                formattedPrice\n                formattedComparePrice\n                formattedDiscountedPrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n                pricePerUnitData @include(if: $withPricePerUnit) {\n                baseQuantity\n                baseMeasurementUnit\n          }\n          seoTitle\n          seoDescription\n          createVersion\n          digitalProductFileItems {\n                fileId\n                fileType\n                fileName\n          }\n          productItems {\n                price\n                comparePrice\n                formattedPrice\n                formattedComparePrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n                optionsSelections\n                isVisible\n                inventory {\n                status\n                quantity\n            }\n            sku\n            weight\n            surcharge\n            subscriptionPlans {\n                list {\n                id\n                price\n                formattedPrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n              }\n            }\n          }\n          name\n          isTrackingInventory\n          inventory {\n            status\n            quantity\n          }\n          isVisible\n          isManageProductItems\n          isInStock\n          media {\n            id\n            url\n            fullUrl\n            altText\n            thumbnailFullUrl: fullUrl(width: 50, height: 50)\n            mediaType\n            videoType\n            videoFiles {\n                url\n                width\n                height\n                format\n                quality\n            }\n            width\n            height\n            index\n            title\n          }\n          customTextFields {\n            title\n            isMandatory\n            inputLimit\n          }\n          nextOptionsSelectionId\n          options {\n            title\n            optionType\n            selections {\n                id\n                value\n                description\n                linkedMediaItems {\n                    altText\n                    url\n                    fullUrl\n                    thumbnailFullUrl: fullUrl(width: 50, height: 50)\n                    mediaType\n                    width\n                    height\n                    index\n                    title\n                    videoFiles {\n                        url\n                        width\n                        height\n                        format\n                        quality\n                    }\n                }\n            }\n          }\n          productType\n          urlPart\n          additionalInfo {\n                id\n            title\n            description\n            index\n          }\n          subscriptionPlans {\n                list(onlyVisible: true) {\n                  id\n              name\n              tagline\n              frequency\n              duration\n              price\n              formattedPrice\n              pricePerUnit @include(if: $withPricePerUnit)\n              formattedPricePerUnit @include(if: $withPricePerUnit)\n            }\n            oneTimePurchase {\n                  index\n            }\n          }\n          discount {\n                mode\n            value\n          }\n          currency\n          weight\n          seoJson\n        }\n      }\n      localeData(language: \"en\") @include(if: $withCountryCodes) {\n            countries {\n              key\n          shortKey\n        }\n      }\n    }", "variables": fmt.Sprintf(`{"slug":"%s","externalId":"","withPricePerUnit":true,"withCountryCodes":false}}`, productID)}

	jsonValue, err := json.Marshal(jsonQuery)

	if err != nil {
		log.Println("Could't marshall json", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://www.lafiancee.com.br/_api/wix-ecommerce-storefront-web/api", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", authorizationToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)

	if err != nil {
		log.Println("Error getting response", err)
		return nil, err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Error reading the data", err)
		return nil, err
	}

	return data, nil
}

func buildProduct(data []byte) Product {
	var pr ProductResult

	json.Unmarshal(data, &pr)

	var product Product
	product.ID = ""
	product.Title = pr.Data.Catalog.Product.Title
	product.Subtitle = ""
	product.ContentURL =
		"https://www.lafiancee.com.br/product-page/" + pr.Data.Catalog.Product.URLPart

	// populating media
	for _, element := range pr.Data.Catalog.Product.Media {
		product.Media = append(product.Media, element.FullUrl)
	}

	// populating the map
	for _, element := range pr.Data.Catalog.Product.Options {
		for _, value := range element.Selections {
			m := make(map[string]string)
			m[element.Title] = value.Value
			product.Attributes = append(product.Attributes, m)
		}
	}

	return product
}

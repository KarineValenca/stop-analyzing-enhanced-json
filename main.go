package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	//productID := "allure-3110"
	jsonQuery := map[string]string{
		"query": "query getProductBySlug($externalId: String!, $slug: String!, $withPricePerUnit: Boolean!, $withCountryCodes: Boolean!) {\n          appSettings(externalId: $externalId) {\n            widgetSettings\n      }\n      catalog {\n            product(slug: $slug, onlyVisible: true) {\n                id\n                description\n                isVisible\n                sku\n                ribbon\n                price\n                comparePrice\n                discountedPrice\n                formattedPrice\n                formattedComparePrice\n                formattedDiscountedPrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n                pricePerUnitData @include(if: $withPricePerUnit) {\n                baseQuantity\n                baseMeasurementUnit\n          }\n          seoTitle\n          seoDescription\n          createVersion\n          digitalProductFileItems {\n                fileId\n                fileType\n                fileName\n          }\n          productItems {\n                price\n                comparePrice\n                formattedPrice\n                formattedComparePrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n                optionsSelections\n                isVisible\n                inventory {\n                status\n                quantity\n            }\n            sku\n            weight\n            surcharge\n            subscriptionPlans {\n                list {\n                id\n                price\n                formattedPrice\n                pricePerUnit @include(if: $withPricePerUnit)\n                formattedPricePerUnit @include(if: $withPricePerUnit)\n              }\n            }\n          }\n          name\n          isTrackingInventory\n          inventory {\n            status\n            quantity\n          }\n          isVisible\n          isManageProductItems\n          isInStock\n          media {\n            id\n            url\n            fullUrl\n            altText\n            thumbnailFullUrl: fullUrl(width: 50, height: 50)\n            mediaType\n            videoType\n            videoFiles {\n                url\n                width\n                height\n                format\n                quality\n            }\n            width\n            height\n            index\n            title\n          }\n          customTextFields {\n            title\n            isMandatory\n            inputLimit\n          }\n          nextOptionsSelectionId\n          options {\n            title\n            optionType\n            selections {\n                id\n                value\n                description\n                linkedMediaItems {\n                    altText\n                    url\n                    fullUrl\n                    thumbnailFullUrl: fullUrl(width: 50, height: 50)\n                    mediaType\n                    width\n                    height\n                    index\n                    title\n                    videoFiles {\n                        url\n                        width\n                        height\n                        format\n                        quality\n                    }\n                }\n            }\n          }\n          productType\n          urlPart\n          additionalInfo {\n                id\n            title\n            description\n            index\n          }\n          subscriptionPlans {\n                list(onlyVisible: true) {\n                  id\n              name\n              tagline\n              frequency\n              duration\n              price\n              formattedPrice\n              pricePerUnit @include(if: $withPricePerUnit)\n              formattedPricePerUnit @include(if: $withPricePerUnit)\n            }\n            oneTimePurchase {\n                  index\n            }\n          }\n          discount {\n                mode\n            value\n          }\n          currency\n          weight\n          seoJson\n        }\n      }\n      localeData(language: \"en\") @include(if: $withCountryCodes) {\n            countries {\n              key\n          shortKey\n        }\n      }\n    }", "variables": `{"slug":"allure-3110","externalId":"","withPricePerUnit":true,"withCountryCodes":false}}`}

	authorizationToken := "brUTfgwc9eaqQ4m_KjbIkjnR-MRt9rGfCLGikGEPiRU.eyJpbnN0YW5jZUlkIjoiMWI0OTQ1ODItZDg5Zi00MmY2LTg0YzAtNTAxOGE3NzI1Y2MyIiwiYXBwRGVmSWQiOiIxMzgwYjcwMy1jZTgxLWZmMDUtZjExNS0zOTU3MWQ5NGRmY2QiLCJtZXRhU2l0ZUlkIjoiN2RlM2ExNjgtNDEyNC00NDljLTg4ZDYtZmViNjkzYWY3NzRjIiwic2lnbkRhdGUiOiIyMDIwLTA5LTIzVDEyOjI3OjE4LjUyOVoiLCJ2ZW5kb3JQcm9kdWN0SWQiOiJQcmVtaXVtMSIsImRlbW9Nb2RlIjpmYWxzZSwiYWlkIjoiOWE0ZjJjNDAtMTIzNC00ZGM3LTg3OWEtMjIzZDMxMzI0N2E1IiwiYmlUb2tlbiI6IjY2YWFlNGVhLTk5YmItMDY2YS0wYzE2LWFlYWUzNGRkMmI4ZSIsInNpdGVPd25lcklkIjoiZmI0Y2Y2ODQtODZkZS00N2E0LWE2NjUtZjE4ZDcxYzA3YzUxIn0"

	data, _ := fetchData(jsonQuery, authorizationToken)
	fmt.Println(string(data))
	fmt.Println(getUrlPart())
}

func fetchData(jsonQuery map[string]string, authorizationToken string) ([]byte, error) {
	jsonValue, err := json.Marshal(jsonQuery)

	if err != nil {
		fmt.Println("Could't marshall json", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://www.lafiancee.com.br/_api/wix-ecommerce-storefront-web/api", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", authorizationToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error getting response", err)
		return nil, err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading the data", err)
		return nil, err
	}

	return data, nil
}

func getUrlPart() ([]string, error) {
	jsonFile, err := os.Open("lafiancee.json")
	if err != nil {
		fmt.Println("Could't open the json file")
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//var result map[string]interface{}
	var result Result
	json.Unmarshal([]byte(byteValue), &result)

	productList := result.Data.Catalog.Category.ProductsWithMetadata.List

	var URLPartArr []string
	for _, element := range productList {
		URLPartArr = append(URLPartArr, element.URLPart)
	}

	return URLPartArr, nil
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//productID := "allure-3110"
	jsonQuery := map[string]string{
		"query": `query getProductBySlug($externalId: String!, $slug: String!, $withPricePerUnit: Boolean!, $withCountryCodes: Boolean!) {
			appSettings(externalId: $externalId) {
				widgetSettings
	}
	catalog {
				product(slug: $slug, onlyVisible: true) {
						id
						description
						isVisible
						sku
						ribbon
						price
						comparePrice
						discountedPrice
						formattedPrice
						formattedComparePrice
						formattedDiscountedPrice
						pricePerUnit @include(if: $withPricePerUnit)
						formattedPricePerUnit @include(if: $withPricePerUnit)
						pricePerUnitData @include(if: $withPricePerUnit) {
						baseQuantity
						baseMeasurementUnit
			}
			seoTitle
			seoDescription
			createVersion
			digitalProductFileItems {
						fileId
						fileType
						fileName
			}
			productItems {
						price
						comparePrice
						formattedPrice
						formattedComparePrice
						pricePerUnit @include(if: $withPricePerUnit)
						formattedPricePerUnit @include(if: $withPricePerUnit)
						optionsSelections
						isVisible
						inventory {
						status
						quantity
				}
				sku
				weight
				surcharge
				subscriptionPlans {
						list {
						id
						price
						formattedPrice
						pricePerUnit @include(if: $withPricePerUnit)
						formattedPricePerUnit @include(if: $withPricePerUnit)
					}
				}
			}
			name
			isTrackingInventory
			inventory {
				status
				quantity
			}
			isVisible
			isManageProductItems
			isInStock
			media {
				id
				url
				fullUrl
				altText
				thumbnailFullUrl: fullUrl(width: 50, height: 50)
				mediaType
				videoType
				videoFiles {
						url
						width
						height
						format
						quality
				}
				width
				height
				index
				title
			}
			customTextFields {
				title
				isMandatory
				inputLimit
			}
			nextOptionsSelectionId
			options {
				title
				optionType
				selections {
						id
						value
						description
						linkedMediaItems {
								altText
								url
								fullUrl
								thumbnailFullUrl: fullUrl(width: 50, height: 50)
								mediaType
								width
								height
								index
								title
								videoFiles {
										url
										width
										height
										format
										quality
								}
						}
				}
			}
			productType
			urlPart
			additionalInfo {
						id
				title
				description
				index
			}
			subscriptionPlans {
						list(onlyVisible: true) {
							id
					name
					tagline
					frequency
					duration
					price
					formattedPrice
					pricePerUnit @include(if: $withPricePerUnit)
					formattedPricePerUnit @include(if: $withPricePerUnit)
				}
				oneTimePurchase {
							index
				}
			}
			discount {
						mode
				value
			}
			currency
			weight
			seoJson
		}
	}
	localeData(language: "en") @include(if: $withCountryCodes) {
				countries {
					key
			shortKey
		}
	}
},"variables":{"slug":"allure-3110","externalId":"","withPricePerUnit":true,"withCountryCodes":false}`,
	}

	jsonValue, err := json.Marshal(jsonQuery)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "https://www.lafiancee.com.br/_api/wix-ecommerce-storefront-web/api", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "brUTfgwc9eaqQ4m_KjbIkjnR-MRt9rGfCLGikGEPiRU.eyJpbnN0YW5jZUlkIjoiMWI0OTQ1ODItZDg5Zi00MmY2LTg0YzAtNTAxOGE3NzI1Y2MyIiwiYXBwRGVmSWQiOiIxMzgwYjcwMy1jZTgxLWZmMDUtZjExNS0zOTU3MWQ5NGRmY2QiLCJtZXRhU2l0ZUlkIjoiN2RlM2ExNjgtNDEyNC00NDljLTg4ZDYtZmViNjkzYWY3NzRjIiwic2lnbkRhdGUiOiIyMDIwLTA5LTIzVDEyOjI3OjE4LjUyOVoiLCJ2ZW5kb3JQcm9kdWN0SWQiOiJQcmVtaXVtMSIsImRlbW9Nb2RlIjpmYWxzZSwiYWlkIjoiOWE0ZjJjNDAtMTIzNC00ZGM3LTg3OWEtMjIzZDMxMzI0N2E1IiwiYmlUb2tlbiI6IjY2YWFlNGVhLTk5YmItMDY2YS0wYzE2LWFlYWUzNGRkMmI4ZSIsInNpdGVPd25lcklkIjoiZmI0Y2Y2ODQtODZkZS00N2E0LWE2NjUtZjE4ZDcxYzA3YzUxIn0")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

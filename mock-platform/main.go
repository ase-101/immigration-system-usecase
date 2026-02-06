// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/v1/otpmanager/otp/generate", func(w http.ResponseWriter, r *http.Request) {
		otp := 111111 // Generate a number between 100000 and 999999
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"response": {"otp": "%d"}}`, otp)
	})

	http.HandleFunc("/v1/auditmanager/audits", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("audit endpoint invoked\n")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "success"}`)
	})

	http.HandleFunc("/v1/notifier/sms/send", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Mock sending OTP\n")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"response": {"status": "success"}}`)
	})

	http.HandleFunc("/auth/realms/mosip/protocol/openid-connect/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("openid-connect/token invoked\n")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token": "test-access-token"}`)
	})

	http.HandleFunc("/masterdata/ui-spec", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"response": {"schema": [
            {
                "id": "email",
                "controlType": "textbox",
                "labelName": {
                    "en": "Email"
                },
                "placeholder": {
                    "eng": "Enter your email"
                },
                "validators": [],
                "required": false,
                "disabled": true,
                "prefix": [],
                "alignmentGroup": "groupA"
            },
            {
                "id": "homeCountry",
                "controlType": "dropdown",
                "labelName": {
                    "eng": "Country"
                },
                "placeholder": {
                    "en": "Select Country"
                },
                "validators": [],
                "alignmentGroup": "groupH",
                "required": true,
				"subType" : "country"
            },
            {
                "id": "passportId",
                "controlType": "textbox",
                "labelName": {
                    "eng": "Passport ID"
                },
                "placeholder": {
                    "en": "Enter Passport ID"
                },
                "required": true,
                "alignmentGroup": "groupM"
            },
			{
                "id": "expireDate",
                "controlType": "date",
                "labelName": {
                    "eng": "Passport Expire Date"
                },
                "alignmentGroup": "expiredate",
                "required": true
            },
            {
                "id": "encodedPhoto",
                "controlType": "photo",
                "labelName": {
                    "eng": "Capture Photo",
                    "khm": "ថតរូប"
                },
                "placeholder": {
                    "eng": "Click to capture photo",
                    "khm": "ចុចដើម្បីថតរូប"
                },
                "info": {
                    "eng": "Please click here to capture your photo using your device's camera.",
                    "khm": "សូមចុចទីនេះដើម្បីថតរូបរបស់អ្នកដោយប្រើកាមេរ៉ារបស់ឧបករណ៍របស់អ្នក។"
                },
                "acceptedFileTypes": "image/jpeg , image/jpg , image/png , image/webp",
                "required": true,
                "alignmentGroup": "groupF"
            },
            {
                "id": "consent",
                "controlType": "checkbox",
                "labelName": {
                    "eng": "I agree to Veridonia’s <b><a href='#'>Terms & Conditions</a></b> and <b><a href='#'>Privacy Policy</a></b>, to store & process my information as required."
                },
                "required": true,
                "alignmentGroup": "groupO"
            }
        ],
        "allowedValues": {
			"country" : { "westalis" : { "eng" : "Westalis"}, "india" : { "eng" : "India"}, "morocco" : { "eng" : "Morocco"}}
		},
        "i18nValues": {
            "errors": {
                "required": {
                    "en": "This field is required"
                }
            },
            "labels": {
                "capturePhoto": {
                    "en": "Capture Photo",
                    "km": "ថតរូប"
                },
                "clickToUpload": {
                    "en": "Click to upload",
                    "km": "ចុចដើម្បីបញ្ចូលឬថតរូប"
                }
            },
            "placeholders": {
            }
        },
        "language": {
            "mandatory": [
                "eng",
				"fra"
            ],
            "optional": [],
            "langCodeMap": {
                "eng": "en",
				"fra":"fr"
            }
        },
        "maxUploadFileSize": 5242880
    }}`)
	})

	http.HandleFunc("/masterdata/identity-schema", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
						"$schema": "https://json-schema.org/draft/2020-12/schema",
						"type": "object",
						"$defs": {
							"langField": {
							"type": "array",
							"items": {
								"type": "object",
								"properties": {
								"language": {
									"type": "string"
								},
								"value": {
									"type": "string"
								}
								},
								"required": [
								"language",
								"value"
								],
								"additionalProperties": false
							}
							}
						},
						"properties": {
							"individualId": {
							"type": "string",
							"pattern": "\\S"
							},
							"passportId": {
							"type": "string"
							},
							"expireDate": {
							"type": "string"
							},
							"consent": {
							"type": "boolean"
							},
							"homeCountry": {
							"type": "string"
							},
							"encodedPhoto": {
							"type": "string"
							}
						},
						"required": [
							"individualId",
							"homeCountry",
							"passportId",
							"email",
							"consent",
							"encodedPhoto"
						],
						"additionalProperties": true
						}`)
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

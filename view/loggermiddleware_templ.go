// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LoggerMiddleware() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"es\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Redirección con Fetch</title><style>\n                body {\n                    font-family: Arial, sans-serif;\n                    text-align: center;\n                }\n                h1 {\n                    font-size: 2rem;\n                    margin-bottom: 1rem;\n                }\n                img {\n                    width: 200px;\n                    margin-top: 1rem;\n                }\n            </style></head><body><h1>Redirigiendo...</h1><img src=\"https://sailingcharter.s3.eu-west-3.amazonaws.com/logo/Logo+Sailing+Paradise+Cuadrado+Negro.png\" alt=\"Sailing Paradise Logo\"><script>\n(async () => {\n    try {\n        // Función para obtener la ubicación del usuario\n        const getLocation = () => {\n            return new Promise((resolve, reject) => {\n                if (navigator.geolocation) {\n                    navigator.geolocation.getCurrentPosition(\n                        (position) => {\n                            const latitud = position.coords.latitude;\n                            const longitud = position.coords.longitude;\n                            resolve({ latitud, longitud });\n                        },\n                        (error) => {\n                            reject(error);\n                        }\n                    );\n                } else {\n                    reject(new Error(\"Geolocalización no es soportada por este navegador.\"));\n                }\n            });\n        };\n\n        // Obtener la ubicación real del usuario\n        const { latitud, longitud } = await getLocation();\n\n        // Obtener hora exacta en formato ISO 8601\n        const fechaActual = new Date().toISOString();\n\n        // Obtener el valor del parámetro \"code\" de la URL actual\n        const urlParams = new URLSearchParams(window.location.search);\n        const code = urlParams.get('qrcode');\n\n        // Construir la URL de la página externa a la que se enviará la información\n        const externalUrl = '/api/visitor'; // Reemplaza con la URL real\n\n        // Enviar la información usando fetch\n        fetch(externalUrl, {\n            method: 'POST',\n            headers: {\n                'Content-Type': 'application/json'\n            },\n            body: JSON.stringify({\n                place:{\n                    latitude: latitud,\n                    longitude: longitud,\n                },\n                time: fechaActual,\n                qrcode: code\n            })\n        });\n\n        //obtener todos los query parameters menos el qrcode\n        const queryParameters = new URLSearchParams(window.location.search);\n        queryParameters.delete('qrcode');\n\n        // Construir la URL de redirección, manteniendo el path y los query parameters\n        const redirectUrl = `https://sailing-paradise.com${window.location.pathname}?${queryParameters}`;\n\n        // Redirigir al usuario\n        window.location.href = redirectUrl;\n    } catch (error) {\n        console.error('Error durante el proceso:', error);\n    }\n})();\n\n            </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

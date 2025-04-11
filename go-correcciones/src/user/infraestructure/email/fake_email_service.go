package email

import "fmt"

func SendWelcomeEmail(email string, username string) {
    fmt.Printf("Simulando envío de correo a %s\n¡Hola %s! Bienvenido a la plataforma.\n", email, username)
}

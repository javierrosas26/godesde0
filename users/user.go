package users

import (
	"fmt"
	"time"

	"github.com/javierrosas26/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(15, "Javier", time.Now(), true)
	fmt.Println(u)
}

package ailang

import (
	"testing"

	// "github.com/anakilang-ai/backend/models"
	helper "github.com/anakilang-ai/backend/utils"
	// modules "github.com/anakilang-ai/backend/moduls"
)

// var db = module.MongoConnect("MONGOSTRING", "ailang")

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := helper.GenerateKey()
	t.Logf("PrivateKey : %v", privateKey)
	t.Logf("PublicKey : %v", publicKey)
}

// // TestInsertOneDoc
// func TestInsertOneDoc(t *testing.T) {
// 	var data = map[string]interface{}{
// 		"username": "teeamai",
// 		"password": "12345",
// 	}
// 	insertedDoc, err := helper.InsertOneDoc(modules.Mongoconn, "users", data)
// 	if err != nil {
// 		t.Errorf("Error : %v", err)
// 	}
// 	t.Logf("InsertedDoc : %v", insertedDoc)
// }

// func TestSignUp(t *testing.T) {
// 	var doc model.User
// 	doc.NamaLengkap = "UserName"
// 	doc.Email = "user@example.com"
// 	doc.Password = "Password123"
// 	doc.Confirmpassword = "Password123"
// 	email, err := module.SignUp(db, "user", doc)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan email:", email)
// 	}
// }

// func TestLogIn(t *testing.T) {
// 	var user models.User
// 	user.Email = "user@example.com"
// 	user.Password = "Password123"
// 	user, err := modules.LogIn(db, user)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Berhasil LogIn : ", user.Email)
// 	}
// }



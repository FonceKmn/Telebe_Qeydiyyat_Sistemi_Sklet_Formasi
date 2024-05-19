package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "strconv"
    "time"
)

type Student struct {
    Id              int
    Ad              string
    Soyad           string
    Kontakt         string
    Steam           string
    QT              string
    Grade1          *int
    Grade2          *int
    Grade3          *int
    Grade4          *int
    FG              *int
}

var db []Student

func init() {
    file, err := ioutil.ReadFile("Data.json")
    if err == nil {
        json.Unmarshal(file, &db)
    }
}

func saveDB() {
    data, _ := json.MarshalIndent(db, "", "  ")
    ioutil.WriteFile("Data.json", data, 0644)
}

func yoxlama(i int) bool {
    for _, student := range db {
        if student.Id == i {
            return true
        }
    }
    return false
}

func Qebul() {
    var ad, soyad, contact, steam string
    fmt.Print("Adi daxil edin: ")
    fmt.Scanln(&ad)
    fmt.Print("Soyadi Daxil edin: ")
    fmt.Scanln(&soyad)
    fmt.Print("Kontakt melumatlarini daxil edin: ")
    fmt.Scanln(&contact)
    fmt.Print("Hansi kursa qeydiyyat olunur ?: ")
    fmt.Scanln(&steam)

    if ad == "" || soyad == "" || contact == "" || steam == "" {
        fmt.Println("Bosh aana elave olunmushdur diqqetli olun")
        Qebul()
        return
    }

    id_rand_6 := rand.Intn(1000000)
    for yoxlama(id_rand_6) {
        id_rand_6 = rand.Intn(1000000)
    }

    qeydiyyat_tarixi := time.Now().Format("2006-01-02 15:04:05")
    newStudent := Student{
        Id:      id_rand_6,
        Ad:      ad,
        Soyad:   soyad,
        Kontakt: contact,
        Steam:   steam,
        QT:      qeydiyyat_tarixi,
    }
    db = append(db, newStudent)
    saveDB()
}

func Goster() {
    for _, item := range db {
        fmt.Println(item)
    }
}

func Axtar() {
    var i int
    fmt.Print("Axtarish ucun Telebenin Id- nomresini qeyd edin: ")
    fmt.Scan(&i)
    for _, student := range db {
        if student.Id == i {
            fmt.Println(student)
            return
        }
    }
    fmt.Println("Student not found")
}

func Silme() {
    var i int
    fmt.Print("Silmek ucun Telebenin Id- nomresini qeyd edin: ")
    fmt.Scan(&i)
    for index, student := range db {
        if student.Id == i {
            fmt.Printf("Silinen Telebe: %v\n", student)
            db = append(db[:index], db[index+1:]...)
            saveDB()
            return
        }
    }
    fmt.Println("Student not found")
}

func Yenile() {
    var i, c int
    var l string
    var grade int

    fmt.Print("Yenilemek Ucun Telebe Id sini daxil edin: ")
    fmt.Scan(&i)

    for index, student := range db {
        if student.Id == i {
            fmt.Printf("Melumatlari Yenilenecek Telebe: %v\n", student)
            fmt.Println(" 1. Ad \n 2. Soyad \n 3. Kontact \n 4. Kurs \n 5. Davamiyyet \n 6. Tapshiriq \n 7. Modul Imtahani \n 8. Final Imtahan")
            fmt.Print("Yenilemek istediyiniz melumat novunu secin: ")
            fmt.Scan(&c)
            switch c {
            case 1:
                fmt.Print("Telebenin yeni Adini daxil edin: ")
                fmt.Scan(&l)
                db[index].Ad = l
            case 2:
                fmt.Print("Telebenin yeni Soyadini daxil edin: ")
                fmt.Scan(&l)
                db[index].Soyad = l
            case 3:
                fmt.Print("Telebenin yeni elaqe novunu qeyd edin: ")
                fmt.Scan(&l)
                db[index].Kontakt = l
            case 4:
                fmt.Print("Telebenin yeni Kursunu daxil edin: ")
                fmt.Scan(&l)
                db[index].Steam = l
            case 5:
                fmt.Print("Telebenin yeni Davamiyyet qiymetini daxil edin: ")
                fmt.Scan(&grade)
                db[index].Grade1 = &grade
            case 6:
                fmt.Print("Telebenin yeni Tapshiriq qiymetini daxil edin: ")
                fmt.Scan(&grade)
                db[index].Grade2 = &grade
            case 7:
                fmt.Print("Telebenin yeni Modul Imtahani neticesini daxil edin: ")
                fmt.Scan(&grade)
                db[index].Grade3 = &grade
            case 8:
                fmt.Print("Telebenin yeni Final Imtahan neticesini daxil edin: ")
                fmt.Scan(&grade)
                db[index].Grade4 = &grade
            default:
                fmt.Println("Invalid selection")
            }
            saveDB()
            return
        }
    }
    fmt.Println("Student not found")
}

func main() {
    rand.Seed(time.Now().UnixNano())
    for {
        fmt.Println("Student management System")
        fmt.Println(" 1. Qebul \n 2. Goster \n 3. Axtar \n 4. Silme \n 5. Yenile \n 6. Cixish")
        var c int
        fmt.Print("Icra olunmaq ucun uygun emelliyati secin: ")
        fmt.Scan(&c)

        switch c {
        case 1:
            fmt.Println("Qebul secenek secilmishdir")
            Qebul()
        case 2:
            fmt.Println("Goster secenek secilmishdir")
            Goster()
        case 3:
            fmt.Println("Axtar secilmishdir")
            Axtar()
        case 4:
            fmt.Println("Silme secilmishdir")
            Silme()
        case 5:
            fmt.Println("Yenile secilmishdir")
            Yenile()
        case 6:
            fmt.Println("Cixish Secilib")
            return
        default:
            fmt.Println("Xeta bash verdi")
        }
    }
}

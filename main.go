package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql","root:@tcp(167.205.67.251:3306)/kos_kosan_di_bandung")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Kos struct {
		Id         		   		int
		Nama	   		   		string
		Jenis	   		   		string
		Biaya   				int
		Fasilitas_Kamar	   			string
		Fasilitas_Kamar_Mandi			string
		Fasilitas_Umum				string
		Akses_Lingkungan			string
		Area					string
	}
	router := gin.Default()

	// GET handler untuk menampilkan data kos sesuai dengan area yang dipilih yang diurutkan berdasarkan biaya kos per bulan (rendah-tinggi)
	router.GET("/kos/:area", func(c *gin.Context) {
		var (
			kos Kos
			semuakos []Kos
		)
		area := c.Param("area")
		rows, err := db.Query("select * from kos where area = ? order by biaya_per_bulan;", area)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&kos.Id, &kos.Nama, &kos.Jenis, &kos.Biaya, &kos.Fasilitas_Kamar, &kos.Fasilitas_Kamar_Mandi, &kos.Fasilitas_Umum, &kos.Akses_Lingkungan, &kos.Area)
			semuakos = append(semuakos, kos)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": semuakos,
			"count":  len(semuakos),
		})
	})

	// GET handler untuk menampilkan data kos putra sesuai dengan area yang dipilih yang diurutkan berdasarkan biaya kos per bulan (rendah-tinggi)
	router.GET("/kosputra/:area", func(c *gin.Context) {
		var (
			kos Kos
			semuakos []Kos
		)
		area := c.Param("area")
		rows, err := db.Query("select * from kos where area = ? and jenis = 'putra' order by biaya_per_bulan;", area)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&kos.Id, &kos.Nama, &kos.Jenis, &kos.Biaya, &kos.Fasilitas_Kamar, &kos.Fasilitas_Kamar_Mandi, &kos.Fasilitas_Umum, &kos.Akses_Lingkungan, &kos.Area)
			semuakos = append(semuakos, kos)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": semuakos,
			"count":  len(semuakos),
		})
	})

	// GET handler untuk menampilkan data kos putri sesuai dengan area yang dipilih yang diurutkan berdasarkan biaya kos per bulan (rendah-tinggi)
	router.GET("/kosputri/:area", func(c *gin.Context) {
		var (
			kos Kos
			semuakos []Kos
		)
		area := c.Param("area")
		rows, err := db.Query("select * from kos where area = ? and jenis = 'putri' order by biaya_per_bulan;", area)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&kos.Id, &kos.Nama, &kos.Jenis, &kos.Biaya, &kos.Fasilitas_Kamar, &kos.Fasilitas_Kamar_Mandi, &kos.Fasilitas_Umum, &kos.Akses_Lingkungan, &kos.Area)
			semuakos = append(semuakos, kos)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": semuakos,
			"count":  len(semuakos),
		})
	})

	// GET handler untuk menampilkan data kos campur sesuai dengan area yang dipilih yang diurutkan berdasarkan biaya kos per bulan (rendah-tinggi)
	router.GET("/koscampur/:area", func(c *gin.Context) {
		var (
			kos Kos
			semuakos []Kos
		)
		area := c.Param("area")
		rows, err := db.Query("select * from kos where area = ? and jenis = 'campur' order by biaya_per_bulan;", area)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&kos.Id, &kos.Nama, &kos.Jenis, &kos.Biaya, &kos.Fasilitas_Kamar, &kos.Fasilitas_Kamar_Mandi, &kos.Fasilitas_Umum, &kos.Akses_Lingkungan, &kos.Area)
			semuakos = append(semuakos, kos)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": semuakos,
			"count":  len(semuakos),
		})
	})

	// GET handler untuk menampilkan seluruh data kos yang ada di database
	router.GET("/kos", func(c *gin.Context) {
		var (
			kos  	 Kos
			semuakos []Kos
		)
		rows, err := db.Query("select * from kos;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&kos.Id, &kos.Nama, &kos.Jenis, &kos.Biaya, &kos.Fasilitas_Kamar, &kos.Fasilitas_Kamar_Mandi, &kos.Fasilitas_Umum, &kos.Akses_Lingkungan, &kos.Area)
			semuakos = append(semuakos, kos)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": semuakos,
			"count":  len(semuakos),
		})
	})

	// POST handler untuk menambahkan data kos baru ke database dengan menggunakan form
	router.POST("/insertkos", func(c *gin.Context) {
		nama := c.PostForm("nama")
		jenis := c.PostForm("jenis")
		biaya_per_bulan := c.PostForm("biaya_per_bulan")
		fasilitas_kamar := c.PostForm("fasilitas_kamar")
		fasilitas_kamar_mandi := c.PostForm("fasilitas_kamar_mandi")
		fasilitas_umum := c.PostForm("fasilitas_umum")
		akses_lingkungan := c.PostForm("akses_lingkungan")
		area := c.PostForm("area")
		stmt, err := db.Prepare("insert into kos (nama, jenis, biaya_per_bulan, fasilitas_kamar, fasilitas_kamar_mandi, fasilitas_umum, akses_lingkungan, area) values(?,?,?,?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nama, jenis, biaya_per_bulan, fasilitas_kamar, fasilitas_kamar_mandi, fasilitas_umum, akses_lingkungan, area)

		if err != nil {
			fmt.Print(err.Error())
		}
		defer stmt.Close()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Data kos berhasil ditambahkan."),
		})
	})

	// PUT handler untuk memperbarui data kos yang sudah ada di database dengan menggunakan form
	router.PUT("/editkos", func(c *gin.Context) {
		id := c.Query("id")
		nama := c.PostForm("nama")
		jenis := c.PostForm("jenis")
		biaya_per_bulan := c.PostForm("biaya_per_bulan")
		fasilitas_kamar := c.PostForm("fasilitas_kamar")
		fasilitas_kamar_mandi := c.PostForm("fasilitas_kamar_mandi")
		fasilitas_umum := c.PostForm("fasilitas_umum")
		akses_lingkungan := c.PostForm("akses_lingkungan")
		area := c.PostForm("area")
		stmt, err := db.Prepare("update kos set nama= ?, jenis= ?, biaya_per_bulan= ?, fasilitas_kamar= ?, fasilitas_kamar_mandi= ?, fasilitas_umum= ?, akses_lingkungan= ?, area= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nama, jenis, biaya_per_bulan, fasilitas_kamar, fasilitas_kamar_mandi, fasilitas_umum, akses_lingkungan, area, id)
		if err != nil {
			fmt.Print(err.Error())
		}
		defer stmt.Close()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Pembaruan data kos berhasil."),
		})
	})

	// DELETE handler untuk menghapus data kos dari database
	router.DELETE("/deletekos", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from kos where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Data kos dengan id %s berhasil dihapus.", id),
		})
	})
	router.Run(":3000")
}

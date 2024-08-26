package Groupie_tracker

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type ArtistWithLocation struct {
	JsonData interface{}
}

var (
	tmpl   *template.Template
	errors AllMessageErrors
)

// Initialize the global template variable
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	errors = ErrorsMessage()
}

func GetDataFromJson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErrors(w, errors.MethodNotAllowed, errors.DescriptionMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}
	var buf bytes.Buffer
	artisData, errs := GetArtistsDataStruct()
	if errs != nil {
		HandleErrors(w, errors.BadRequest, errors.DescriptionBadRequest, http.StatusBadRequest)
		return
	}

	err := tmpl.ExecuteTemplate(&buf, "index.html", artisData)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	_, erro := buf.WriteTo(w)
	if erro != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
}

func HandlerShowRelation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErrors(w, errors.MethodNotAllowed, errors.DescriptionMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	// if r.Method == "POST" {
    //     city := r.FormValue("city")
    //     fmt.Println(city) 
    // }
	idParam := r.PathValue("id")
	cities := r.PostFormValue("city")
	fmt.Println(cities)
	artist, err := FetchDataRelationFromId(idParam, cities)
	if err != nil {
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}

	if artist.Id == 0 {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	errs := tmpl.ExecuteTemplate(&buf, "InforArtis.html", artist)
	if errs != nil {
		fmt.Println(errs)
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
	_, errs = buf.WriteTo(w)
	if errs != nil {
		fmt.Println(errs)
		HandleErrors(w, errors.InternalError, errors.DescriptionInternalError, http.StatusInternalServerError)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	var lat, lon float64
	var err error
	if city != "" {
		lat, lon, err = GetCoordinates(city)
		if err != nil {
			lat, lon = 0, 0
		}
	}
	data := SendData(lat, lon, city)
	tmpl.ExecuteTemplate(w, "maps.html", data)
}

func HandleStyle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
			return
		}
	}()
	path := r.URL.Path[len("/styles"):]
	fullpath := filepath.Join("src", path)
	fileinfo, err := os.Stat(fullpath)

	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.StripPrefix("/styles", http.FileServer(http.Dir("src"))).ServeHTTP(w, r)
	} else {
		HandleErrors(w, errors.NotFound, errors.DescriptionNotFound, http.StatusNotFound)
		return
	}
}

func HandleErrors(w http.ResponseWriter, message, description string, code int) {
	errorsMessage := Errors{
		Message:     message,
		Description: description,
		Code:        code,
	}
	w.WriteHeader(code)
	err := tmpl.ExecuteTemplate(w, "errors.html", errorsMessage)
	if err != nil {
		http.Error(w, "Error 500 Internal Server Error", http.StatusInternalServerError)
	}
}

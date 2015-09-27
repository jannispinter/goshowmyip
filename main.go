// goshowmyip is a simple IP address and User Agent displaying web app written
// in Go. As an additional gimmick, the web app fetches random kitten pictures
// from flickr.

// goshowmyip is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// goshowmyip is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.

package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var (
	error_url      = "data:image/png;base64,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"
	api_key        = ""
	api_search_url = "https://api.flickr.com/services/rest/?api_key=" + api_key + "&method=flickr.photos.search&tags=kitten&safe_search=1&license=1,2,3,4,5,6,7&format=json&nojsoncallback=1"
)

// helper function to make api calls
func fetchApi(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func fetchPhotoDetails(photoId string) (license, user string, err error) {
	api_photo_url := "https://api.flickr.com/services/rest/?api_key=" + api_key + "&method=flickr.photos.getInfo&photo_id=" + photoId + "&format=json&nojsoncallback=1"
	response, err := fetchApi(api_photo_url)
	if err != nil {
		return "", "", err
	}

	var photo PhotoResult
	json.Unmarshal(response, &photo)

	license = getLicense(photo.Photo.License)
	user = photo.Photo.Owner.Username

	return license, user, err

}

// fetches a random kitten image and some additional meta data from flickr
func getRandomKittenPicture() (url, license, user, photoid string) {
	response, err := fetchApi(api_search_url)
	if err != nil {
		return error_url, "", "", ""
	}

	var result SearchResult
	json.Unmarshal(response, &result)
	randomPhoto := result.Photos.Photo[rand.Intn(len(result.Photos.Photo))]
	url = fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_z.jpg", randomPhoto.Farm, randomPhoto.Server, randomPhoto.Id, randomPhoto.Secret)

	license, user, err = fetchPhotoDetails(randomPhoto.Id)
	if err != nil {
		return error_url, "", "", ""
	}

	return url, license, user, randomPhoto.Id
}

// helper function to convert a license id into human readable text
// See: https://www.flickr.com/services/api/flickr.photos.licenses.getInfo.html
func getLicense(id string) string {

	switch id {
	case "1":
		return "CC-BY-NC-SA"
	case "2":
		return "CC-BY-NC"
	case "3":
		return "CC-BY-NC-ND"
	case "4":
		return "CC-BY"
	case "5":
		return "CC-BY-SA"
	case "6":
		return "CC-BY-ND"
	default:
		return "No known copyright restrictions"
	}

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	template := "<html><head><title>Whats my IP and UA</title></head><body><center><h1>Your IP address is:</h1>%s<h3>Your user agent is:</h3>%s<h3>Your random kitten picture:</h3><img src=\"%s\"/><br><small>%s <a href=\"https://www.flickr.com/photos/%s/%s/\">%s</a></small></center></body></html>"

	url, license, user, photoid := getRandomKittenPicture()
	fmt.Fprintf(w, template, r.Header.Get("X-Real-IP"), r.UserAgent(), url, license, user, photoid, user)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page 投稿ページ
type Page struct {
	Title string
	Body  []byte
}

// ポインタレシーバ
// Write処理, 戻り値はerror, txtファイルにBodyの内容を書き込む
func (page *Page) save() error {

	filename := page.Title + ".txt"
	// (対象ファイル, 書き込む内容, 権限)
	return ioutil.WriteFile(filename, page.Body, 0600)
}

// Read処理(戻り値は２つ)
func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	// Bodyなし
	if err != nil {
		return nil, err
	}
	// Bodyあり
	return &Page{
		Title: title,
		Body:  body}, nil
}

// レンダリング処理
func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, page)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	// エラーハンドリング(何もなければeditにリダイレクト)
	if err != nil {
		http.Redirect(w, r, "edit"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// title取得
	title := r.URL.Path[len("/save/"):]
	// body取得
	body := r.FormValue("body")
	// データ保存
	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	// エラーハンドリング
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// webサーバーの立ち上げ：(取得パス, 動かす関数)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	// エラーが出たら強制終了
	log.Fatal(http.ListenAndServe(":8000", nil))
}

package main

type Data struct {
	Menu DataMenu `json:"menu"`
}

type DataMenu struct {
	ID    string        `json:"id"`
	Popup DataMenuPopup `json:"popup"`
	Value string        `json:"value"`
}

type DataMenuPopup struct {
	Menuitem []DataMenuPopupMenuitem `json:"menuitem"`
}

type DataMenuPopupMenuitem struct {
	Onclick string `json:"onclick"`
	Value   string `json:"value"`
}
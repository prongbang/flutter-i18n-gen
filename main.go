package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	appname := flag.String("appname", "appName", "")
	source := flag.String("source", "i18n/en.json", "")
	target := flag.String("target", "lib/localization/key_localizations.dart", "")
	flag.Parse()

	data, err := ioutil.ReadFile(*source)
	if err != nil {
		log.Println("File reading error", err)
		return
	}

	maps := map[string]string{}
	if err := json.Unmarshal(data, &maps); err != nil {
		log.Println("Unmarshal error", err)
	}

	keyLocalize := `
	import 'dart:ui';

	import 'package:%s/localization/base_localizations.dart';

	class KeyLocalizations extends BaseLocalizations {
	  KeyLocalizations(Locale locale) : super(locale);
	  %s
	}`
	localize := ""
	for key := range maps {
		localize = localize + fmt.Sprintf("\n    String get %s => translate(\"%s\");", key, key)
	}

	if err := ioutil.WriteFile(*target, []byte(fmt.Sprintf(keyLocalize, *appname, localize)), 0644); err != nil {
		log.Println("Generate file error", err)
	} else {
		log.Println(fmt.Sprintf("Generate file success"))
	}
}

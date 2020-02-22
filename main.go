package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

		keys := strings.Split(key, "_")
		keyName := ""
		for i := 0; i < len(keys); i++ {
			if i == 0 {
				keyName = keys[i]
				continue
			}
			first := strings.ToUpper(keys[i][0:1])
			second := keys[i][1:]
			keyName = keyName + first + second
		}

		localize = localize + fmt.Sprintf("\n    String get %s => translate(\"%s\");", keyName, key)
	}

	if err := ioutil.WriteFile(*target, []byte(fmt.Sprintf(keyLocalize, *appname, localize)), 0644); err != nil {
		log.Println("Generate file error", err)
	} else {
		log.Println(fmt.Sprintf("Generate file success"))
	}
}

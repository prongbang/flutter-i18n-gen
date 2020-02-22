import 'dart:ui';

import 'package:appName/localization/base_localizations.dart';

class KeyLocalizations extends BaseLocalizations {
  KeyLocalizations(Locale locale) : super(locale);
  
  String get firstSecond => translate("first_second");
  String get firstSecondLast => translate("first_second_last");
  String get title => translate("title");
}
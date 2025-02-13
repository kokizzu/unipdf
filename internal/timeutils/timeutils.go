//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package timeutils ;import (_f "errors";_ff "fmt";_dc "regexp";_fb "strconv";_b "time";);func FormatPdfTime (in _b .Time )string {_g :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_df ,_ :=_fb .ParseInt (_g [1:3],10,32);_db ,_ :=_fb .ParseInt (_g [4:6],10,32);
_ffb :=int64 (in .Year ());_bd :=int64 (in .Month ());_ffe :=int64 (in .Day ());_ba :=int64 (in .Hour ());_af :=int64 (in .Minute ());_fd :=int64 (in .Second ());_c :=_g [0];return _ff .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_ffb ,_bd ,_ffe ,_ba ,_af ,_fd ,_c ,_df ,_db );
};var _bf =_dc .MustCompile ("\u005cs\u002a\u0044\u005cs\u002a\u003a\u005cs\u002a(\\\u0064\u007b\u0034\u007d\u0029\u0028\u005cd\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064{2\u007d)\u003f\u0028\u005b\u002b\u002d\u005a]\u0029\u003f\u0028\u005c\u0064{\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032}\u0029\u003f");
func ParsePdfTime (pdfTime string )(_b .Time ,error ){_dcb :=_bf .FindAllStringSubmatch (pdfTime ,1);if len (_dcb )< 1{if len (pdfTime )> 0&&pdfTime [0]!='D'{pdfTime =_ff .Sprintf ("\u0044\u003a\u0025\u0073",pdfTime );return ParsePdfTime (pdfTime );};return _b .Time {},_ff .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );
};if len (_dcb [0])!=10{return _b .Time {},_f .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_fc ,_ :=_fb .ParseInt (_dcb [0][1],10,32);_bdf ,_ :=_fb .ParseInt (_dcb [0][2],10,32);_bab ,_ :=_fb .ParseInt (_dcb [0][3],10,32);_gca ,_ :=_fb .ParseInt (_dcb [0][4],10,32);_fg ,_ :=_fb .ParseInt (_dcb [0][5],10,32);_dg ,_ :=_fb .ParseInt (_dcb [0][6],10,32);
var (_e byte ;_ge int64 ;_dd int64 ;);_e ='+';if len (_dcb [0][7])> 0{if _dcb [0][7]=="\u002d"{_e ='-';}else if _dcb [0][7]=="\u005a"{_e ='Z';};};if len (_dcb [0][8])> 0{_ge ,_ =_fb .ParseInt (_dcb [0][8],10,32);}else {_ge =0;};if len (_dcb [0][9])> 0{_dd ,_ =_fb .ParseInt (_dcb [0][9],10,32);
}else {_dd =0;};_ca :=int (_ge *60*60+_dd *60);switch _e {case '-':_ca =-_ca ;case 'Z':_ca =0;};_cad :=_ff .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_e ,_ge ,_dd );_faf :=_b .FixedZone (_cad ,_ca );return _b .Date (int (_fc ),_b .Month (_bdf ),int (_bab ),int (_gca ),int (_fg ),int (_dg ),0,_faf ),nil ;
};
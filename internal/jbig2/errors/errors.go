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

package errors ;import (_b "fmt";_fg "golang.org/x/xerrors";);func Wrap (err error ,processName ,message string )error {if _bg ,_gab :=err .(*processError );_gab {_bg ._a ="";};_dd :=_ga (message ,processName );_dd ._d =err ;return _dd ;};func _ga (_ea ,_eg string )*processError {return &processError {_a :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_g :_ea ,_e :_eg };
};type processError struct{_a string ;_e string ;_g string ;_d error ;};var _ _fg .Wrapper =(*processError )(nil );func (_fb *processError )Unwrap ()error {return _fb ._d };func (_ge *processError )Error ()string {var _af string ;if _ge ._a !=""{_af =_ge ._a ;
};_af +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_ge ._e ;if _ge ._g !=""{_af +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_ge ._g ;};if _ge ._d !=nil {_af +="\u002e\u0020"+_ge ._d .Error ();};return _af ;};func Error (processName ,message string )error {return _ga (message ,processName )};
func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _gd ,_gde :=err .(*processError );_gde {_gd ._a ="";};_gf :=_ga (_b .Sprintf (message ,arguments ...),processName );_gf ._d =err ;return _gf ;};func Errorf (processName ,message string ,arguments ...interface{})error {return _ga (_b .Sprintf (message ,arguments ...),processName );
};
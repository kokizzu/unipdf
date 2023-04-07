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

package basic ;import _e "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_fc *NumSlice )Add (v float32 ){*_fc =append (*_fc ,v )};func (_ef NumSlice )GetIntSlice ()[]int {_ec :=make ([]int ,len (_ef ));
for _bg ,_feaa :=range _ef {_ec [_bg ]=int (_feaa );};return _ec ;};func (_f IntsMap )Add (key uint64 ,value int ){_f [key ]=append (_f [key ],value )};func (_cg *Stack )Peek ()(_ag interface{},_aa bool ){return _cg .peek ()};func (_cb *IntSlice )Copy ()*IntSlice {_bf :=IntSlice (make ([]int ,len (*_cb )));
copy (_bf ,*_cb );return &_bf ;};func NewIntSlice (i int )*IntSlice {_ff :=IntSlice (make ([]int ,i ));return &_ff };func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;
};func (_cbe *Stack )peek ()(interface{},bool ){_bfc :=_cbe .top ();if _bfc ==-1{return nil ,false ;};return _cbe .Data [_bfc ],true ;};type Stack struct{Data []interface{};Aux *Stack ;};func (_dd IntSlice )Size ()int {return len (_dd )};func (_bd NumSlice )GetInt (i int )(int ,error ){const _fcd ="\u0047\u0065\u0074\u0049\u006e\u0074";
if i < 0||i > len (_bd )-1{return 0,_e .Errorf (_fcd ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_eg :=_bd [i ];return int (_eg +Sign (_eg )*0.5),nil ;};
func Min (x ,y int )int {if x < y {return x ;};return y ;};type IntSlice []int ;type IntsMap map[uint64 ][]int ;func (_cc *NumSlice )AddInt (v int ){*_cc =append (*_cc ,float32 (v ))};type NumSlice []float32 ;func (_g IntsMap )GetSlice (key uint64 )([]int ,bool ){_a ,_gb :=_g [key ];
if !_gb {return nil ,false ;};return _a ,true ;};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_b IntsMap )Delete (key uint64 ){delete (_b ,key )};func (_fe IntsMap )Get (key uint64 )(int ,bool ){_ca ,_fea :=_fe [key ];if !_fea {return 0,false ;
};if len (_ca )==0{return 0,false ;};return _ca [0],true ;};func (_ba *Stack )Len ()int {return len (_ba .Data )};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_d *IntSlice )Add (v int )error {if _d ==nil {return _e .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_d =append (*_d ,v );return nil ;};func (_fd *Stack )top ()int {return len (_fd .Data )-1};func (_ae NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_ae )-1{return 0,_e .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _ae [i ],nil ;};func (_bad *Stack )Pop ()(_ge interface{},_ce bool ){_ge ,_ce =_bad .peek ();if !_ce {return nil ,_ce ;};_bad .Data =_bad .Data [:_bad .top ()];return _ge ,true ;};func NewNumSlice (i int )*NumSlice {_cd :=NumSlice (make ([]float32 ,i ));
return &_cd };func (_cab IntSlice )Get (index int )(int ,error ){if index > len (_cab )-1{return 0,_e .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _cab [index ],nil ;};func (_aaf *Stack )Push (v interface{}){_aaf .Data =append (_aaf .Data ,v )};
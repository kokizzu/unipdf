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

package basic ;import _c "github.com/unidoc/unipdf/v3/internal/jbig2/errors";type IntsMap map[uint64 ][]int ;func NewIntSlice (i int )*IntSlice {_ga :=IntSlice (make ([]int ,i ));return &_ga };type NumSlice []float32 ;func (_aa *NumSlice )Add (v float32 ){*_aa =append (*_aa ,v )};
func (_aga *Stack )Push (v interface{}){_aga .Data =append (_aga .Data ,v )};type IntSlice []int ;func (_ag NumSlice )GetInt (i int )(int ,error ){const _b ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_ag )-1{return 0,_c .Errorf (_b ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};_gge :=_ag [i ];return int (_gge +Sign (_gge )*0.5),nil ;};func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_e IntSlice )Size ()int {return len (_e )};
type Stack struct{Data []interface{};Aux *Stack ;};func (_gac *IntSlice )Add (v int )error {if _gac ==nil {return _c .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_gac =append (*_gac ,v );return nil ;};func (_dac NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_dac )-1{return 0,_c .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _dac [i ],nil ;};func (_fda IntsMap )Delete (key uint64 ){delete (_fda ,key )};func (_db NumSlice )GetIntSlice ()[]int {_gacf :=make ([]int ,len (_db ));for _ae ,_ef :=range _db {_gacf [_ae ]=int (_ef );};return _gacf ;};func (_d *IntSlice )Copy ()*IntSlice {_ac :=IntSlice (make ([]int ,len (*_d )));
copy (_ac ,*_d );return &_ac ;};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_da IntSlice )Get (index int )(int ,error ){if index > len (_da )-1{return 0,_c .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _da [index ],nil ;};func (_ab *Stack )Pop ()(_ee interface{},_dec bool ){_ee ,_dec =_ab .peek ();if !_dec {return nil ,_dec ;};_ab .Data =_ab .Data [:_ab .top ()];return _ee ,true ;};func (_fa *NumSlice )AddInt (v int ){*_fa =append (*_fa ,float32 (v ))};
func (_gea *Stack )Peek ()(_be interface{},_ff bool ){return _gea .peek ()};func (_gf IntsMap )Add (key uint64 ,value int ){_gf [key ]=append (_gf [key ],value )};func (_ffa *Stack )top ()int {return len (_ffa .Data )-1};func Max (x ,y int )int {if x > y {return x ;
};return y ;};func (_ge *Stack )Len ()int {return len (_ge .Data )};func Min (x ,y int )int {if x < y {return x ;};return y ;};func (_f IntsMap )Get (key uint64 )(int ,bool ){_a ,_fd :=_f [key ];if !_fd {return 0,false ;};if len (_a )==0{return 0,false ;
};return _a [0],true ;};func (_ca IntsMap )GetSlice (key uint64 )([]int ,bool ){_gg ,_fb :=_ca [key ];if !_fb {return nil ,false ;};return _gg ,true ;};func (_bea *Stack )peek ()(interface{},bool ){_bd :=_bea .top ();if _bd ==-1{return nil ,false ;};return _bea .Data [_bd ],true ;
};func NewNumSlice (i int )*NumSlice {_de :=NumSlice (make ([]float32 ,i ));return &_de };func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};
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

package sampling ;import (_d "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/imageutil";_a "io";);func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _eb []uint32 ;_ee :=bitsPerSample ;var _bb uint32 ;
var _bf byte ;_gc :=0;_ab :=0;_ce :=0;for _ce < len (data ){if _gc > 0{_fa :=_gc ;if _ee < _fa {_fa =_ee ;};_bb =(_bb <<uint (_fa ))|uint32 (_bf >>uint (8-_fa ));_gc -=_fa ;if _gc > 0{_bf =_bf <<uint (_fa );}else {_bf =0;};_ee -=_fa ;if _ee ==0{_eb =append (_eb ,_bb );
_ee =bitsPerSample ;_bb =0;_ab ++;};}else {_gf :=data [_ce ];_ce ++;_cc :=8;if _ee < _cc {_cc =_ee ;};_gc =8-_cc ;_bb =(_bb <<uint (_cc ))|uint32 (_gf >>uint (_gc ));if _cc < 8{_bf =_gf <<uint (_cc );};_ee -=_cc ;if _ee ==0{_eb =append (_eb ,_bb );_ee =bitsPerSample ;
_bb =0;_ab ++;};};};for _gc >=bitsPerSample {_gff :=_gc ;if _ee < _gff {_gff =_ee ;};_bb =(_bb <<uint (_gff ))|uint32 (_bf >>uint (8-_gff ));_gc -=_gff ;if _gc > 0{_bf =_bf <<uint (_gff );}else {_bf =0;};_ee -=_gff ;if _ee ==0{_eb =append (_eb ,_bb );_ee =bitsPerSample ;
_bb =0;_ab ++;};};return _eb ;};func (_ea *Reader )ReadSample ()(uint32 ,error ){if _ea ._afe ==_ea ._e .Height {return 0,_a .EOF ;};_bed ,_gb :=_ea ._de .ReadBits (byte (_ea ._e .BitsPerComponent ));if _gb !=nil {return 0,_gb ;};_ea ._ff --;if _ea ._ff ==0{_ea ._ff =_ea ._e .ColorComponents ;
_ea ._af ++;};if _ea ._af ==_ea ._e .Width {if _ea ._be {_ea ._de .ConsumeRemainingBits ();};_ea ._af =0;_ea ._afe ++;};return uint32 (_bed ),nil ;};type SampleWriter interface{WriteSample (_ac uint32 )error ;WriteSamples (_gd []uint32 )error ;};type SampleReader interface{ReadSample ()(uint32 ,error );
ReadSamples (_f []uint32 )error ;};type Reader struct{_e _b .ImageBase ;_de *_d .Reader ;_af ,_afe ,_ff int ;_be bool ;};func (_fad *Writer )WriteSamples (samples []uint32 )error {for _dbe :=0;_dbe < len (samples );_dbe ++{if _fb :=_fad .WriteSample (samples [_dbe ]);
_fb !=nil {return _fb ;};};return nil ;};func NewReader (img _b .ImageBase )*Reader {return &Reader {_de :_d .NewReader (img .Data ),_e :img ,_ff :img .ColorComponents ,_be :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func (_c *Reader )ReadSamples (samples []uint32 )(_gbf error ){for _cb :=0;_cb < len (samples );_cb ++{samples [_cb ],_gbf =_c .ReadSample ();if _gbf !=nil {return _gbf ;};};return nil ;};func NewWriter (img _b .ImageBase )*Writer {return &Writer {_ecc :_d .NewWriterMSB (img .Data ),_ec :img ,_bfe :img .ColorComponents ,_bd :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func (_cf *Writer )WriteSample (sample uint32 )error {if _ ,_fac :=_cf ._ecc .WriteBits (uint64 (sample ),_cf ._ec .BitsPerComponent );_fac !=nil {return _fac ;};_cf ._bfe --;if _cf ._bfe ==0{_cf ._bfe =_cf ._ec .ColorComponents ;_cf ._cd ++;};if _cf ._cd ==_cf ._ec .Width {if _cf ._bd {_cf ._ecc .FinishByte ();
};_cf ._cd =0;};return nil ;};type Writer struct{_ec _b .ImageBase ;_ecc *_d .Writer ;_cd ,_bfe int ;_bd bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _gg []uint32 ;_beb :=bitsPerOutputSample ;
var _ga uint32 ;var _bc uint32 ;_dc :=0;_faa :=0;_ced :=0;for _ced < len (data ){if _dc > 0{_dg :=_dc ;if _beb < _dg {_dg =_beb ;};_ga =(_ga <<uint (_dg ))|(_bc >>uint (bitsPerInputSample -_dg ));_dc -=_dg ;if _dc > 0{_bc =_bc <<uint (_dg );}else {_bc =0;
};_beb -=_dg ;if _beb ==0{_gg =append (_gg ,_ga );_beb =bitsPerOutputSample ;_ga =0;_faa ++;};}else {_aff :=data [_ced ];_ced ++;_dcd :=bitsPerInputSample ;if _beb < _dcd {_dcd =_beb ;};_dc =bitsPerInputSample -_dcd ;_ga =(_ga <<uint (_dcd ))|(_aff >>uint (_dc ));
if _dcd < bitsPerInputSample {_bc =_aff <<uint (_dcd );};_beb -=_dcd ;if _beb ==0{_gg =append (_gg ,_ga );_beb =bitsPerOutputSample ;_ga =0;_faa ++;};};};for _dc >=bitsPerOutputSample {_db :=_dc ;if _beb < _db {_db =_beb ;};_ga =(_ga <<uint (_db ))|(_bc >>uint (bitsPerInputSample -_db ));
_dc -=_db ;if _dc > 0{_bc =_bc <<uint (_db );}else {_bc =0;};_beb -=_db ;if _beb ==0{_gg =append (_gg ,_ga );_beb =bitsPerOutputSample ;_ga =0;_faa ++;};};if _beb > 0&&_beb < bitsPerOutputSample {_ga <<=uint (_beb );_gg =append (_gg ,_ga );};return _gg ;
};
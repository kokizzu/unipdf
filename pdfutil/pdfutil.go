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

package pdfutil ;import (_b "github.com/unidoc/unipdf/v3/common";_ef "github.com/unidoc/unipdf/v3/contentstream";_d "github.com/unidoc/unipdf/v3/contentstream/draw";_g "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_e .PdfPage )error {_c ,_bg :=page .GetMediaBox ();if _bg !=nil {return _bg ;};_f ,_bg :=page .GetRotate ();if _bg !=nil {_b .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_bg .Error ());
};_fg :=_f %360!=0&&_f %90==0;_c .Normalize ();_de ,_eb ,_gc ,_ae :=_c .Llx ,_c .Lly ,_c .Width (),_c .Height ();_ce :=_de !=0||_eb !=0;if !_fg &&!_ce {return nil ;};_ac :=func (_gg ,_cc ,_gf float64 )_d .BoundingBox {return _d .Path {Points :[]_d .Point {_d .NewPoint (0,0).Rotate (_gf ),_d .NewPoint (_gg ,0).Rotate (_gf ),_d .NewPoint (0,_cc ).Rotate (_gf ),_d .NewPoint (_gg ,_cc ).Rotate (_gf )}}.GetBoundingBox ();
};_ge :=_ef .NewContentCreator ();var _bgc float64 ;if _fg {_bgc =-float64 (_f );_fd :=_ac (_gc ,_ae ,_bgc );_ge .Translate ((_fd .Width -_gc )/2+_gc /2,(_fd .Height -_ae )/2+_ae /2);_ge .RotateDeg (_bgc );_ge .Translate (-_gc /2,-_ae /2);_gc ,_ae =_fd .Width ,_fd .Height ;
};if _ce {_ge .Translate (-_de ,-_eb );};_fge :=_ge .Operations ();_eba ,_bg :=_g .MakeStream (_fge .Bytes (),_g .NewFlateEncoder ());if _bg !=nil {return _bg ;};_bb :=_g .MakeArray (_eba );_bb .Append (page .GetContentStreamObjs ()...);*_c =_e .PdfRectangle {Urx :_gc ,Ury :_ae };
if _cg :=page .CropBox ;_cg !=nil {_cg .Normalize ();_aa ,_ed ,_cgf ,_ec :=_cg .Llx -_de ,_cg .Lly -_eb ,_cg .Width (),_cg .Height ();if _fg {_dd :=_ac (_cgf ,_ec ,_bgc );_cgf ,_ec =_dd .Width ,_dd .Height ;};*_cg =_e .PdfRectangle {Llx :_aa ,Lly :_ed ,Urx :_aa +_cgf ,Ury :_ed +_ec };
};_b .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_bgc ,_fge ,_c );page .Contents =_bb ;page .Rotate =nil ;
return nil ;};
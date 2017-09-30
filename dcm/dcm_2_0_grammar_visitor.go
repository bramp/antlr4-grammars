// Generated from DCM_2_0_grammar.g4 by ANTLR 4.7.

package dcm // DCM_2_0_grammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DCM_2_0_grammarParser.
type DCM_2_0_grammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DCM_2_0_grammarParser#konservierung.
	VisitKonservierung(ctx *KonservierungContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kons_kopf.
	VisitKons_kopf(ctx *Kons_kopfContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#modulkopf_info.
	VisitModulkopf_info(ctx *Modulkopf_infoContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#mod_zeile.
	VisitMod_zeile(ctx *Mod_zeileContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#mod_anf_zeile.
	VisitMod_anf_zeile(ctx *Mod_anf_zeileContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#mod_fort_zeile.
	VisitMod_fort_zeile(ctx *Mod_fort_zeileContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#mod_ele_name.
	VisitMod_ele_name(ctx *Mod_ele_nameContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#mod_ele_wert.
	VisitMod_ele_wert(ctx *Mod_ele_wertContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#funktionsdef.
	VisitFunktionsdef(ctx *FunktionsdefContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#funktionszeile.
	VisitFunktionszeile(ctx *FunktionszeileContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#fkt_version.
	VisitFkt_version(ctx *Fkt_versionContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#fkt_langname.
	VisitFkt_langname(ctx *Fkt_langnameContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#variantendef.
	VisitVariantendef(ctx *VariantendefContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#variantenkrit.
	VisitVariantenkrit(ctx *VariantenkritContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#krit_name.
	VisitKrit_name(ctx *Krit_nameContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#krit_wert.
	VisitKrit_wert(ctx *Krit_wertContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kons_rumpf.
	VisitKons_rumpf(ctx *Kons_rumpfContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kenngroesse.
	VisitKenngroesse(ctx *KenngroesseContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kennwert.
	VisitKennwert(ctx *KennwertContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kennwerteblock.
	VisitKennwerteblock(ctx *KennwerteblockContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kennlinie.
	VisitKennlinie(ctx *KennlinieContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kennfeld.
	VisitKennfeld(ctx *KennfeldContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#gruppenstuetzstellen.
	VisitGruppenstuetzstellen(ctx *GruppenstuetzstellenContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kenntext.
	VisitKenntext(ctx *KenntextContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kgr_info.
	VisitKgr_info(ctx *Kgr_infoContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#einheit_x.
	VisitEinheit_x(ctx *Einheit_xContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#einheit_y.
	VisitEinheit_y(ctx *Einheit_yContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#einheit_w.
	VisitEinheit_w(ctx *Einheit_wContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#langname.
	VisitLangname(ctx *LangnameContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#displayname.
	VisitDisplayname(ctx *DisplaynameContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#var_abhangigkeiten.
	VisitVar_abhangigkeiten(ctx *Var_abhangigkeitenContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#var_abh.
	VisitVar_abh(ctx *Var_abhContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#funktionszugehorigkeit.
	VisitFunktionszugehorigkeit(ctx *FunktionszugehorigkeitContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#anzahl_x.
	VisitAnzahl_x(ctx *Anzahl_xContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#anzahl_y.
	VisitAnzahl_y(ctx *Anzahl_yContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#werteliste.
	VisitWerteliste(ctx *WertelisteContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#werteliste_kwb.
	VisitWerteliste_kwb(ctx *Werteliste_kwbContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#sst_liste_x.
	VisitSst_liste_x(ctx *Sst_liste_xContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kf_zeile_liste.
	VisitKf_zeile_liste(ctx *Kf_zeile_listeContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kf_zeile_liste_r.
	VisitKf_zeile_liste_r(ctx *Kf_zeile_liste_rContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#kf_zeile_liste_tx.
	VisitKf_zeile_liste_tx(ctx *Kf_zeile_liste_txContext) interface{}

	// Visit a parse tree produced by DCM_2_0_grammarParser#realzahl.
	VisitRealzahl(ctx *RealzahlContext) interface{}
}

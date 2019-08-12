// Code generated from DCM_2_0_grammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dcm_2_0_grammar // DCM_2_0_grammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDCM_2_0_grammarListener is a complete listener for a parse tree produced by DCM_2_0_grammarParser.
type BaseDCM_2_0_grammarListener struct{}

var _ DCM_2_0_grammarListener = &BaseDCM_2_0_grammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDCM_2_0_grammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDCM_2_0_grammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDCM_2_0_grammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDCM_2_0_grammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKonservierung is called when production konservierung is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKonservierung(ctx *KonservierungContext) {}

// ExitKonservierung is called when production konservierung is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKonservierung(ctx *KonservierungContext) {}

// EnterKons_kopf is called when production kons_kopf is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKons_kopf(ctx *Kons_kopfContext) {}

// ExitKons_kopf is called when production kons_kopf is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKons_kopf(ctx *Kons_kopfContext) {}

// EnterModulkopf_info is called when production modulkopf_info is entered.
func (s *BaseDCM_2_0_grammarListener) EnterModulkopf_info(ctx *Modulkopf_infoContext) {}

// ExitModulkopf_info is called when production modulkopf_info is exited.
func (s *BaseDCM_2_0_grammarListener) ExitModulkopf_info(ctx *Modulkopf_infoContext) {}

// EnterMod_zeile is called when production mod_zeile is entered.
func (s *BaseDCM_2_0_grammarListener) EnterMod_zeile(ctx *Mod_zeileContext) {}

// ExitMod_zeile is called when production mod_zeile is exited.
func (s *BaseDCM_2_0_grammarListener) ExitMod_zeile(ctx *Mod_zeileContext) {}

// EnterMod_anf_zeile is called when production mod_anf_zeile is entered.
func (s *BaseDCM_2_0_grammarListener) EnterMod_anf_zeile(ctx *Mod_anf_zeileContext) {}

// ExitMod_anf_zeile is called when production mod_anf_zeile is exited.
func (s *BaseDCM_2_0_grammarListener) ExitMod_anf_zeile(ctx *Mod_anf_zeileContext) {}

// EnterMod_fort_zeile is called when production mod_fort_zeile is entered.
func (s *BaseDCM_2_0_grammarListener) EnterMod_fort_zeile(ctx *Mod_fort_zeileContext) {}

// ExitMod_fort_zeile is called when production mod_fort_zeile is exited.
func (s *BaseDCM_2_0_grammarListener) ExitMod_fort_zeile(ctx *Mod_fort_zeileContext) {}

// EnterMod_ele_name is called when production mod_ele_name is entered.
func (s *BaseDCM_2_0_grammarListener) EnterMod_ele_name(ctx *Mod_ele_nameContext) {}

// ExitMod_ele_name is called when production mod_ele_name is exited.
func (s *BaseDCM_2_0_grammarListener) ExitMod_ele_name(ctx *Mod_ele_nameContext) {}

// EnterMod_ele_wert is called when production mod_ele_wert is entered.
func (s *BaseDCM_2_0_grammarListener) EnterMod_ele_wert(ctx *Mod_ele_wertContext) {}

// ExitMod_ele_wert is called when production mod_ele_wert is exited.
func (s *BaseDCM_2_0_grammarListener) ExitMod_ele_wert(ctx *Mod_ele_wertContext) {}

// EnterFunktionsdef is called when production funktionsdef is entered.
func (s *BaseDCM_2_0_grammarListener) EnterFunktionsdef(ctx *FunktionsdefContext) {}

// ExitFunktionsdef is called when production funktionsdef is exited.
func (s *BaseDCM_2_0_grammarListener) ExitFunktionsdef(ctx *FunktionsdefContext) {}

// EnterFunktionszeile is called when production funktionszeile is entered.
func (s *BaseDCM_2_0_grammarListener) EnterFunktionszeile(ctx *FunktionszeileContext) {}

// ExitFunktionszeile is called when production funktionszeile is exited.
func (s *BaseDCM_2_0_grammarListener) ExitFunktionszeile(ctx *FunktionszeileContext) {}

// EnterFkt_version is called when production fkt_version is entered.
func (s *BaseDCM_2_0_grammarListener) EnterFkt_version(ctx *Fkt_versionContext) {}

// ExitFkt_version is called when production fkt_version is exited.
func (s *BaseDCM_2_0_grammarListener) ExitFkt_version(ctx *Fkt_versionContext) {}

// EnterFkt_langname is called when production fkt_langname is entered.
func (s *BaseDCM_2_0_grammarListener) EnterFkt_langname(ctx *Fkt_langnameContext) {}

// ExitFkt_langname is called when production fkt_langname is exited.
func (s *BaseDCM_2_0_grammarListener) ExitFkt_langname(ctx *Fkt_langnameContext) {}

// EnterVariantendef is called when production variantendef is entered.
func (s *BaseDCM_2_0_grammarListener) EnterVariantendef(ctx *VariantendefContext) {}

// ExitVariantendef is called when production variantendef is exited.
func (s *BaseDCM_2_0_grammarListener) ExitVariantendef(ctx *VariantendefContext) {}

// EnterVariantenkrit is called when production variantenkrit is entered.
func (s *BaseDCM_2_0_grammarListener) EnterVariantenkrit(ctx *VariantenkritContext) {}

// ExitVariantenkrit is called when production variantenkrit is exited.
func (s *BaseDCM_2_0_grammarListener) ExitVariantenkrit(ctx *VariantenkritContext) {}

// EnterKrit_name is called when production krit_name is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKrit_name(ctx *Krit_nameContext) {}

// ExitKrit_name is called when production krit_name is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKrit_name(ctx *Krit_nameContext) {}

// EnterKrit_wert is called when production krit_wert is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKrit_wert(ctx *Krit_wertContext) {}

// ExitKrit_wert is called when production krit_wert is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKrit_wert(ctx *Krit_wertContext) {}

// EnterKons_rumpf is called when production kons_rumpf is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKons_rumpf(ctx *Kons_rumpfContext) {}

// ExitKons_rumpf is called when production kons_rumpf is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKons_rumpf(ctx *Kons_rumpfContext) {}

// EnterKenngroesse is called when production kenngroesse is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKenngroesse(ctx *KenngroesseContext) {}

// ExitKenngroesse is called when production kenngroesse is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKenngroesse(ctx *KenngroesseContext) {}

// EnterKennwert is called when production kennwert is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKennwert(ctx *KennwertContext) {}

// ExitKennwert is called when production kennwert is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKennwert(ctx *KennwertContext) {}

// EnterKennwerteblock is called when production kennwerteblock is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKennwerteblock(ctx *KennwerteblockContext) {}

// ExitKennwerteblock is called when production kennwerteblock is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKennwerteblock(ctx *KennwerteblockContext) {}

// EnterKennlinie is called when production kennlinie is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKennlinie(ctx *KennlinieContext) {}

// ExitKennlinie is called when production kennlinie is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKennlinie(ctx *KennlinieContext) {}

// EnterKennfeld is called when production kennfeld is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKennfeld(ctx *KennfeldContext) {}

// ExitKennfeld is called when production kennfeld is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKennfeld(ctx *KennfeldContext) {}

// EnterGruppenstuetzstellen is called when production gruppenstuetzstellen is entered.
func (s *BaseDCM_2_0_grammarListener) EnterGruppenstuetzstellen(ctx *GruppenstuetzstellenContext) {}

// ExitGruppenstuetzstellen is called when production gruppenstuetzstellen is exited.
func (s *BaseDCM_2_0_grammarListener) ExitGruppenstuetzstellen(ctx *GruppenstuetzstellenContext) {}

// EnterKenntext is called when production kenntext is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKenntext(ctx *KenntextContext) {}

// ExitKenntext is called when production kenntext is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKenntext(ctx *KenntextContext) {}

// EnterKgr_info is called when production kgr_info is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKgr_info(ctx *Kgr_infoContext) {}

// ExitKgr_info is called when production kgr_info is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKgr_info(ctx *Kgr_infoContext) {}

// EnterEinheit_x is called when production einheit_x is entered.
func (s *BaseDCM_2_0_grammarListener) EnterEinheit_x(ctx *Einheit_xContext) {}

// ExitEinheit_x is called when production einheit_x is exited.
func (s *BaseDCM_2_0_grammarListener) ExitEinheit_x(ctx *Einheit_xContext) {}

// EnterEinheit_y is called when production einheit_y is entered.
func (s *BaseDCM_2_0_grammarListener) EnterEinheit_y(ctx *Einheit_yContext) {}

// ExitEinheit_y is called when production einheit_y is exited.
func (s *BaseDCM_2_0_grammarListener) ExitEinheit_y(ctx *Einheit_yContext) {}

// EnterEinheit_w is called when production einheit_w is entered.
func (s *BaseDCM_2_0_grammarListener) EnterEinheit_w(ctx *Einheit_wContext) {}

// ExitEinheit_w is called when production einheit_w is exited.
func (s *BaseDCM_2_0_grammarListener) ExitEinheit_w(ctx *Einheit_wContext) {}

// EnterLangname is called when production langname is entered.
func (s *BaseDCM_2_0_grammarListener) EnterLangname(ctx *LangnameContext) {}

// ExitLangname is called when production langname is exited.
func (s *BaseDCM_2_0_grammarListener) ExitLangname(ctx *LangnameContext) {}

// EnterDisplayname is called when production displayname is entered.
func (s *BaseDCM_2_0_grammarListener) EnterDisplayname(ctx *DisplaynameContext) {}

// ExitDisplayname is called when production displayname is exited.
func (s *BaseDCM_2_0_grammarListener) ExitDisplayname(ctx *DisplaynameContext) {}

// EnterVar_abhangigkeiten is called when production var_abhangigkeiten is entered.
func (s *BaseDCM_2_0_grammarListener) EnterVar_abhangigkeiten(ctx *Var_abhangigkeitenContext) {}

// ExitVar_abhangigkeiten is called when production var_abhangigkeiten is exited.
func (s *BaseDCM_2_0_grammarListener) ExitVar_abhangigkeiten(ctx *Var_abhangigkeitenContext) {}

// EnterVar_abh is called when production var_abh is entered.
func (s *BaseDCM_2_0_grammarListener) EnterVar_abh(ctx *Var_abhContext) {}

// ExitVar_abh is called when production var_abh is exited.
func (s *BaseDCM_2_0_grammarListener) ExitVar_abh(ctx *Var_abhContext) {}

// EnterFunktionszugehorigkeit is called when production funktionszugehorigkeit is entered.
func (s *BaseDCM_2_0_grammarListener) EnterFunktionszugehorigkeit(ctx *FunktionszugehorigkeitContext) {
}

// ExitFunktionszugehorigkeit is called when production funktionszugehorigkeit is exited.
func (s *BaseDCM_2_0_grammarListener) ExitFunktionszugehorigkeit(ctx *FunktionszugehorigkeitContext) {}

// EnterAnzahl_x is called when production anzahl_x is entered.
func (s *BaseDCM_2_0_grammarListener) EnterAnzahl_x(ctx *Anzahl_xContext) {}

// ExitAnzahl_x is called when production anzahl_x is exited.
func (s *BaseDCM_2_0_grammarListener) ExitAnzahl_x(ctx *Anzahl_xContext) {}

// EnterAnzahl_y is called when production anzahl_y is entered.
func (s *BaseDCM_2_0_grammarListener) EnterAnzahl_y(ctx *Anzahl_yContext) {}

// ExitAnzahl_y is called when production anzahl_y is exited.
func (s *BaseDCM_2_0_grammarListener) ExitAnzahl_y(ctx *Anzahl_yContext) {}

// EnterWerteliste is called when production werteliste is entered.
func (s *BaseDCM_2_0_grammarListener) EnterWerteliste(ctx *WertelisteContext) {}

// ExitWerteliste is called when production werteliste is exited.
func (s *BaseDCM_2_0_grammarListener) ExitWerteliste(ctx *WertelisteContext) {}

// EnterWerteliste_kwb is called when production werteliste_kwb is entered.
func (s *BaseDCM_2_0_grammarListener) EnterWerteliste_kwb(ctx *Werteliste_kwbContext) {}

// ExitWerteliste_kwb is called when production werteliste_kwb is exited.
func (s *BaseDCM_2_0_grammarListener) ExitWerteliste_kwb(ctx *Werteliste_kwbContext) {}

// EnterSst_liste_x is called when production sst_liste_x is entered.
func (s *BaseDCM_2_0_grammarListener) EnterSst_liste_x(ctx *Sst_liste_xContext) {}

// ExitSst_liste_x is called when production sst_liste_x is exited.
func (s *BaseDCM_2_0_grammarListener) ExitSst_liste_x(ctx *Sst_liste_xContext) {}

// EnterKf_zeile_liste is called when production kf_zeile_liste is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKf_zeile_liste(ctx *Kf_zeile_listeContext) {}

// ExitKf_zeile_liste is called when production kf_zeile_liste is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKf_zeile_liste(ctx *Kf_zeile_listeContext) {}

// EnterKf_zeile_liste_r is called when production kf_zeile_liste_r is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKf_zeile_liste_r(ctx *Kf_zeile_liste_rContext) {}

// ExitKf_zeile_liste_r is called when production kf_zeile_liste_r is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKf_zeile_liste_r(ctx *Kf_zeile_liste_rContext) {}

// EnterKf_zeile_liste_tx is called when production kf_zeile_liste_tx is entered.
func (s *BaseDCM_2_0_grammarListener) EnterKf_zeile_liste_tx(ctx *Kf_zeile_liste_txContext) {}

// ExitKf_zeile_liste_tx is called when production kf_zeile_liste_tx is exited.
func (s *BaseDCM_2_0_grammarListener) ExitKf_zeile_liste_tx(ctx *Kf_zeile_liste_txContext) {}

// EnterRealzahl is called when production realzahl is entered.
func (s *BaseDCM_2_0_grammarListener) EnterRealzahl(ctx *RealzahlContext) {}

// ExitRealzahl is called when production realzahl is exited.
func (s *BaseDCM_2_0_grammarListener) ExitRealzahl(ctx *RealzahlContext) {}

// Generated from DCM_2_0_grammar.g4 by ANTLR 4.7.

package dcm // DCM_2_0_grammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DCM_2_0_grammarListener is a complete listener for a parse tree produced by DCM_2_0_grammarParser.
type DCM_2_0_grammarListener interface {
	antlr.ParseTreeListener

	// EnterKonservierung is called when entering the konservierung production.
	EnterKonservierung(c *KonservierungContext)

	// EnterKons_kopf is called when entering the kons_kopf production.
	EnterKons_kopf(c *Kons_kopfContext)

	// EnterModulkopf_info is called when entering the modulkopf_info production.
	EnterModulkopf_info(c *Modulkopf_infoContext)

	// EnterMod_zeile is called when entering the mod_zeile production.
	EnterMod_zeile(c *Mod_zeileContext)

	// EnterMod_anf_zeile is called when entering the mod_anf_zeile production.
	EnterMod_anf_zeile(c *Mod_anf_zeileContext)

	// EnterMod_fort_zeile is called when entering the mod_fort_zeile production.
	EnterMod_fort_zeile(c *Mod_fort_zeileContext)

	// EnterMod_ele_name is called when entering the mod_ele_name production.
	EnterMod_ele_name(c *Mod_ele_nameContext)

	// EnterMod_ele_wert is called when entering the mod_ele_wert production.
	EnterMod_ele_wert(c *Mod_ele_wertContext)

	// EnterFunktionsdef is called when entering the funktionsdef production.
	EnterFunktionsdef(c *FunktionsdefContext)

	// EnterFunktionszeile is called when entering the funktionszeile production.
	EnterFunktionszeile(c *FunktionszeileContext)

	// EnterFkt_version is called when entering the fkt_version production.
	EnterFkt_version(c *Fkt_versionContext)

	// EnterFkt_langname is called when entering the fkt_langname production.
	EnterFkt_langname(c *Fkt_langnameContext)

	// EnterVariantendef is called when entering the variantendef production.
	EnterVariantendef(c *VariantendefContext)

	// EnterVariantenkrit is called when entering the variantenkrit production.
	EnterVariantenkrit(c *VariantenkritContext)

	// EnterKrit_name is called when entering the krit_name production.
	EnterKrit_name(c *Krit_nameContext)

	// EnterKrit_wert is called when entering the krit_wert production.
	EnterKrit_wert(c *Krit_wertContext)

	// EnterKons_rumpf is called when entering the kons_rumpf production.
	EnterKons_rumpf(c *Kons_rumpfContext)

	// EnterKenngroesse is called when entering the kenngroesse production.
	EnterKenngroesse(c *KenngroesseContext)

	// EnterKennwert is called when entering the kennwert production.
	EnterKennwert(c *KennwertContext)

	// EnterKennwerteblock is called when entering the kennwerteblock production.
	EnterKennwerteblock(c *KennwerteblockContext)

	// EnterKennlinie is called when entering the kennlinie production.
	EnterKennlinie(c *KennlinieContext)

	// EnterKennfeld is called when entering the kennfeld production.
	EnterKennfeld(c *KennfeldContext)

	// EnterGruppenstuetzstellen is called when entering the gruppenstuetzstellen production.
	EnterGruppenstuetzstellen(c *GruppenstuetzstellenContext)

	// EnterKenntext is called when entering the kenntext production.
	EnterKenntext(c *KenntextContext)

	// EnterKgr_info is called when entering the kgr_info production.
	EnterKgr_info(c *Kgr_infoContext)

	// EnterEinheit_x is called when entering the einheit_x production.
	EnterEinheit_x(c *Einheit_xContext)

	// EnterEinheit_y is called when entering the einheit_y production.
	EnterEinheit_y(c *Einheit_yContext)

	// EnterEinheit_w is called when entering the einheit_w production.
	EnterEinheit_w(c *Einheit_wContext)

	// EnterLangname is called when entering the langname production.
	EnterLangname(c *LangnameContext)

	// EnterDisplayname is called when entering the displayname production.
	EnterDisplayname(c *DisplaynameContext)

	// EnterVar_abhangigkeiten is called when entering the var_abhangigkeiten production.
	EnterVar_abhangigkeiten(c *Var_abhangigkeitenContext)

	// EnterVar_abh is called when entering the var_abh production.
	EnterVar_abh(c *Var_abhContext)

	// EnterFunktionszugehorigkeit is called when entering the funktionszugehorigkeit production.
	EnterFunktionszugehorigkeit(c *FunktionszugehorigkeitContext)

	// EnterAnzahl_x is called when entering the anzahl_x production.
	EnterAnzahl_x(c *Anzahl_xContext)

	// EnterAnzahl_y is called when entering the anzahl_y production.
	EnterAnzahl_y(c *Anzahl_yContext)

	// EnterWerteliste is called when entering the werteliste production.
	EnterWerteliste(c *WertelisteContext)

	// EnterWerteliste_kwb is called when entering the werteliste_kwb production.
	EnterWerteliste_kwb(c *Werteliste_kwbContext)

	// EnterSst_liste_x is called when entering the sst_liste_x production.
	EnterSst_liste_x(c *Sst_liste_xContext)

	// EnterKf_zeile_liste is called when entering the kf_zeile_liste production.
	EnterKf_zeile_liste(c *Kf_zeile_listeContext)

	// EnterKf_zeile_liste_r is called when entering the kf_zeile_liste_r production.
	EnterKf_zeile_liste_r(c *Kf_zeile_liste_rContext)

	// EnterKf_zeile_liste_tx is called when entering the kf_zeile_liste_tx production.
	EnterKf_zeile_liste_tx(c *Kf_zeile_liste_txContext)

	// EnterRealzahl is called when entering the realzahl production.
	EnterRealzahl(c *RealzahlContext)

	// ExitKonservierung is called when exiting the konservierung production.
	ExitKonservierung(c *KonservierungContext)

	// ExitKons_kopf is called when exiting the kons_kopf production.
	ExitKons_kopf(c *Kons_kopfContext)

	// ExitModulkopf_info is called when exiting the modulkopf_info production.
	ExitModulkopf_info(c *Modulkopf_infoContext)

	// ExitMod_zeile is called when exiting the mod_zeile production.
	ExitMod_zeile(c *Mod_zeileContext)

	// ExitMod_anf_zeile is called when exiting the mod_anf_zeile production.
	ExitMod_anf_zeile(c *Mod_anf_zeileContext)

	// ExitMod_fort_zeile is called when exiting the mod_fort_zeile production.
	ExitMod_fort_zeile(c *Mod_fort_zeileContext)

	// ExitMod_ele_name is called when exiting the mod_ele_name production.
	ExitMod_ele_name(c *Mod_ele_nameContext)

	// ExitMod_ele_wert is called when exiting the mod_ele_wert production.
	ExitMod_ele_wert(c *Mod_ele_wertContext)

	// ExitFunktionsdef is called when exiting the funktionsdef production.
	ExitFunktionsdef(c *FunktionsdefContext)

	// ExitFunktionszeile is called when exiting the funktionszeile production.
	ExitFunktionszeile(c *FunktionszeileContext)

	// ExitFkt_version is called when exiting the fkt_version production.
	ExitFkt_version(c *Fkt_versionContext)

	// ExitFkt_langname is called when exiting the fkt_langname production.
	ExitFkt_langname(c *Fkt_langnameContext)

	// ExitVariantendef is called when exiting the variantendef production.
	ExitVariantendef(c *VariantendefContext)

	// ExitVariantenkrit is called when exiting the variantenkrit production.
	ExitVariantenkrit(c *VariantenkritContext)

	// ExitKrit_name is called when exiting the krit_name production.
	ExitKrit_name(c *Krit_nameContext)

	// ExitKrit_wert is called when exiting the krit_wert production.
	ExitKrit_wert(c *Krit_wertContext)

	// ExitKons_rumpf is called when exiting the kons_rumpf production.
	ExitKons_rumpf(c *Kons_rumpfContext)

	// ExitKenngroesse is called when exiting the kenngroesse production.
	ExitKenngroesse(c *KenngroesseContext)

	// ExitKennwert is called when exiting the kennwert production.
	ExitKennwert(c *KennwertContext)

	// ExitKennwerteblock is called when exiting the kennwerteblock production.
	ExitKennwerteblock(c *KennwerteblockContext)

	// ExitKennlinie is called when exiting the kennlinie production.
	ExitKennlinie(c *KennlinieContext)

	// ExitKennfeld is called when exiting the kennfeld production.
	ExitKennfeld(c *KennfeldContext)

	// ExitGruppenstuetzstellen is called when exiting the gruppenstuetzstellen production.
	ExitGruppenstuetzstellen(c *GruppenstuetzstellenContext)

	// ExitKenntext is called when exiting the kenntext production.
	ExitKenntext(c *KenntextContext)

	// ExitKgr_info is called when exiting the kgr_info production.
	ExitKgr_info(c *Kgr_infoContext)

	// ExitEinheit_x is called when exiting the einheit_x production.
	ExitEinheit_x(c *Einheit_xContext)

	// ExitEinheit_y is called when exiting the einheit_y production.
	ExitEinheit_y(c *Einheit_yContext)

	// ExitEinheit_w is called when exiting the einheit_w production.
	ExitEinheit_w(c *Einheit_wContext)

	// ExitLangname is called when exiting the langname production.
	ExitLangname(c *LangnameContext)

	// ExitDisplayname is called when exiting the displayname production.
	ExitDisplayname(c *DisplaynameContext)

	// ExitVar_abhangigkeiten is called when exiting the var_abhangigkeiten production.
	ExitVar_abhangigkeiten(c *Var_abhangigkeitenContext)

	// ExitVar_abh is called when exiting the var_abh production.
	ExitVar_abh(c *Var_abhContext)

	// ExitFunktionszugehorigkeit is called when exiting the funktionszugehorigkeit production.
	ExitFunktionszugehorigkeit(c *FunktionszugehorigkeitContext)

	// ExitAnzahl_x is called when exiting the anzahl_x production.
	ExitAnzahl_x(c *Anzahl_xContext)

	// ExitAnzahl_y is called when exiting the anzahl_y production.
	ExitAnzahl_y(c *Anzahl_yContext)

	// ExitWerteliste is called when exiting the werteliste production.
	ExitWerteliste(c *WertelisteContext)

	// ExitWerteliste_kwb is called when exiting the werteliste_kwb production.
	ExitWerteliste_kwb(c *Werteliste_kwbContext)

	// ExitSst_liste_x is called when exiting the sst_liste_x production.
	ExitSst_liste_x(c *Sst_liste_xContext)

	// ExitKf_zeile_liste is called when exiting the kf_zeile_liste production.
	ExitKf_zeile_liste(c *Kf_zeile_listeContext)

	// ExitKf_zeile_liste_r is called when exiting the kf_zeile_liste_r production.
	ExitKf_zeile_liste_r(c *Kf_zeile_liste_rContext)

	// ExitKf_zeile_liste_tx is called when exiting the kf_zeile_liste_tx production.
	ExitKf_zeile_liste_tx(c *Kf_zeile_liste_txContext)

	// ExitRealzahl is called when exiting the realzahl production.
	ExitRealzahl(c *RealzahlContext)
}

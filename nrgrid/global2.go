package nrgrid

// DmrsSchInfo contains information of PDSCH/PUSCH DMRS per antenna port.
type DmrsSchInfo struct{
	CdmGroup int
	Delta int
}

// refer to 3GPP 38.211 vf30
//  Table 7.4.1.1.2-1: Parameters for PDSCH DM-RS configuration type 1.
//  Table 6.4.1.1.3-1: Parameters for PUSCH DM-RS configuration type 1.
var DmrsSchCfgType1 = map[int]DmrsSchInfo{
	0: {0, 0},
	1: {0, 0},
	2: {1, 1},
	3: {1, 1},
	4: {0, 0},
	5: {0, 0},
	6: {1, 1},
	7: {1, 1},
}

// refer to 3GPP 38.211 vf30
//  Table 7.4.1.1.2-2: Parameters for PDSCH DM-RS configuration type 2.
//  Table 6.4.1.1.3-2: Parameters for PUSCH DM-RS configuration type 2.
var DmrsSchCfgType2 = map[int]DmrsSchInfo{
	0: {0, 0},
	1: {0, 0},
	2: {1, 2},
	3: {1, 2},
	4: {2, 4},
	5: {2, 4},
	6: {0, 0},
	7: {0, 0},
	8: {1, 2},
	9: {1, 2},
	10: {2, 4},
	11: {2, 4},
}

// refer to 3GPP 38.211 vf30
//  Table 7.4.1.1.2-3: PDSCH DM-RS positions l- for single-symbol DM-RS.
// key="td_mapping type_additional position"
//  key = '%s_%s_%s' % (td, self.nrDci11PdschTimeAllocMappingTypeComb.currentText(), self.nrDmrsDedPdschAddPosComb.currentText())
var DmrsPdschPosOneSymb = map[string][]int{
	"2_Type A_pos0": nil, "2_Type A_pos1": nil, "2_Type A_pos2": nil, "2_Type A_pos3": nil,
	"3_Type A_pos0": {0}, "3_Type A_pos1": {0}, "3_Type A_pos2": {0}, "3_Type A_pos3": {0},
	"4_Type A_pos0": {0}, "4_Type A_pos1": {0}, "4_Type A_pos2": {0}, "4_Type A_pos3": {0},
	"5_Type A_pos0": {0}, "5_Type A_pos1": {0}, "5_Type A_pos2": {0}, "5_Type A_pos3": {0},
	"6_Type A_pos0": {0}, "6_Type A_pos1": {0}, "6_Type A_pos2": {0}, "6_Type A_pos3": {0},
	"7_Type A_pos0": {0}, "7_Type A_pos1": {0}, "7_Type A_pos2": {0}, "7_Type A_pos3": {0},
	"8_Type A_pos0": {0}, "8_Type A_pos1": {0, 7}, "8_Type A_pos2": {0, 7}, "8_Type A_pos3": {0, 7},
	"9_Type A_pos0": {0}, "9_Type A_pos1": {0, 7}, "9_Type A_pos2": {0, 7}, "9_Type A_pos3": {0, 7},
	"10_Type A_pos0": {0}, "10_Type A_pos1": {0, 9}, "10_Type A_pos2": {0, 6, 9}, "10_Type A_pos3": {0, 6, 9},
	"11_Type A_pos0": {0}, "11_Type A_pos1": {0, 9}, "11_Type A_pos2": {0, 6, 9}, "11_Type A_pos3": {0, 6, 9},
	"12_Type A_pos0": {0}, "12_Type A_pos1": {0, 9}, "12_Type A_pos2": {0, 6, 9}, "12_Type A_pos3": {0, 5, 8, 11},
	"13_Type A_pos0": {0}, "13_Type A_pos1": {0, 11}, "13_Type A_pos2": {0, 7, 11}, "13_Type A_pos3": {0, 5, 8, 11},
	"14_Type A_pos0": {0}, "14_Type A_pos1": {0, 11}, "14_Type A_pos2": {0, 7, 11}, "14_Type A_pos3": {0, 5, 8, 11},
	"2_Type B_pos0": {0}, "2_Type B_pos1": {0}, "2_Type B_pos2": nil, "2_Type B_pos3": nil,
	"3_Type B_pos0": nil, "3_Type B_pos1": nil, "3_Type B_pos2": nil, "3_Type B_pos3": nil,
	"4_Type B_pos0": {0}, "4_Type B_pos1": {0}, "4_Type B_pos2": nil, "4_Type B_pos3": nil,
	"5_Type B_pos0": nil, "5_Type B_pos1": nil, "5_Type B_pos2": nil, "5_Type B_pos3": nil,
	"6_Type B_pos0": {0}, "6_Type B_pos1": {0, 4}, "6_Type B_pos2": nil, "6_Type B_pos3": nil,
	"7_Type B_pos0": {0}, "7_Type B_pos1": {0, 4}, "7_Type B_pos2": nil, "7_Type B_pos3": nil,
	"8_Type B_pos0": nil, "8_Type B_pos1": nil, "8_Type B_pos2": nil, "8_Type B_pos3": nil,
	"9_Type B_pos0": nil, "9_Type B_pos1": nil, "9_Type B_pos2": nil, "9_Type B_pos3": nil,
	"10_Type B_pos0": nil, "10_Type B_pos1": nil, "10_Type B_pos2": nil, "10_Type B_pos3": nil,
	"11_Type B_pos0": nil, "11_Type B_pos1": nil, "11_Type B_pos2": nil, "11_Type B_pos3": nil,
	"12_Type B_pos0": nil, "12_Type B_pos1": nil, "12_Type B_pos2": nil, "12_Type B_pos3": nil,
	"13_Type B_pos0": nil, "13_Type B_pos1": nil, "13_Type B_pos2": nil, "13_Type B_pos3": nil,
	"14_Type B_pos0": nil, "14_Type B_pos1": nil, "14_Type B_pos2": nil, "14_Type B_pos3": nil,
}

// refer to 3GPP 38.211 vf30
//  Table 7.4.1.1.2-4: PDSCH DM-RS positions l- for double-symbol DM-RS.
// key="td_mapping type_additional position"
//  key = '%s_%s_%s' % (td, self.nrDci11PdschTimeAllocMappingTypeComb.currentText(), self.nrDmrsDedPdschAddPosComb.currentText())
var DmrsPdschPosTwoSymbs = map[string][]int{
	"2_Type A_pos0": nil, "2_Type A_pos1": nil,
	"3_Type A_pos0": nil, "3_Type A_pos1": nil,
	"4_Type A_pos0": {0}, "4_Type A_pos1": {0},
	"5_Type A_pos0": {0}, "5_Type A_pos1": {0},
	"6_Type A_pos0": {0}, "6_Type A_pos1": {0},
	"7_Type A_pos0": {0}, "7_Type A_pos1": {0},
	"8_Type A_pos0": {0}, "8_Type A_pos1": {0},
	"9_Type A_pos0": {0}, "9_Type A_pos1": {0},
	"10_Type A_pos0": {0}, "10_Type A_pos1": {0, 8},
	"11_Type A_pos0": {0}, "11_Type A_pos1": {0, 8},
	"12_Type A_pos0": {0}, "12_Type A_pos1": {0, 8},
	"13_Type A_pos0": {0}, "13_Type A_pos1": {0, 10},
	"14_Type A_pos0": {0}, "14_Type A_pos1": {0, 10},
	"2_Type B_pos0": nil, "2_Type B_pos1": nil,
	"3_Type B_pos0": nil, "3_Type B_pos1": nil,
	"4_Type B_pos0": nil, "4_Type B_pos1": nil,
	"5_Type B_pos0": nil, "5_Type B_pos1": nil,
	"6_Type B_pos0": {0}, "6_Type B_pos1": {0},
	"7_Type B_pos0": {0}, "7_Type B_pos1": {0},
	"8_Type B_pos0": nil, "8_Type B_pos1": nil,
	"9_Type B_pos0": nil, "9_Type B_pos1": nil,
	"10_Type B_pos0": nil, "10_Type B_pos1": nil,
	"11_Type B_pos0": nil, "11_Type B_pos1": nil,
	"12_Type B_pos0": nil, "12_Type B_pos1": nil,
	"13_Type B_pos0": nil, "13_Type B_pos1": nil,
	"14_Type B_pos0": nil, "14_Type B_pos1": nil,
}

// refer to 3GPP 38.211 vf30
//  Table 6.4.1.1.3-3: PUSCH DM-RS positions l- within a slot for single-symbol DM-RS and intra-slot frequency hopping disabled.
// key="ld_mapping type_additional position"
//  key = '%s_%s_%s' % (ld, self.nrDci01PuschTimeAllocMappingTypeComb.currentText(), self.nrDmrsDedPuschAddPosComb.currentText())
var DmrsPuschPosOneSymbWoIntraSlotFh = map[string][]int{
	"1_Type A_pos0": nil, "1_Type A_pos1": nil, "1_Type A_pos2": nil, "1_Type A_pos3": nil,
	"2_Type A_pos0": nil, "2_Type A_pos1": nil, "2_Type A_pos2": nil, "2_Type A_pos3": nil,
	"3_Type A_pos0": nil, "3_Type A_pos1": nil, "3_Type A_pos2": nil, "3_Type A_pos3": nil,
	"4_Type A_pos0": {0}, "4_Type A_pos1": {0}, "4_Type A_pos2": {0}, "4_Type A_pos3": {0},
	"5_Type A_pos0": {0}, "5_Type A_pos1": {0}, "5_Type A_pos2": {0}, "5_Type A_pos3": {0},
	"6_Type A_pos0": {0}, "6_Type A_pos1": {0}, "6_Type A_pos2": {0}, "6_Type A_pos3": {0},
	"7_Type A_pos0": {0}, "7_Type A_pos1": {0}, "7_Type A_pos2": {0}, "7_Type A_pos3": {0},
	"8_Type A_pos0": {0}, "8_Type A_pos1": {0, 7}, "8_Type A_pos2": {0, 7}, "8_Type A_pos3": {0, 7},
	"9_Type A_pos0": {0}, "9_Type A_pos1": {0, 7}, "9_Type A_pos2": {0, 7}, "9_Type A_pos3": {0, 7},
	"10_Type A_pos0": {0}, "10_Type A_pos1": {0, 9}, "10_Type A_pos2": {0, 6, 9}, "10_Type A_pos3": {0, 6, 9},
	"11_Type A_pos0": {0}, "11_Type A_pos1": {0, 9}, "11_Type A_pos2": {0, 6, 9}, "11_Type A_pos3": {0, 6, 9},
	"12_Type A_pos0": {0}, "12_Type A_pos1": {0, 9}, "12_Type A_pos2": {0, 6, 9}, "12_Type A_pos3": {0, 5, 8, 11},
	"13_Type A_pos0": {0}, "13_Type A_pos1": {0, 11}, "13_Type A_pos2": {0, 7, 11}, "13_Type A_pos3": {0, 5, 8, 11},
	"14_Type A_pos0": {0}, "14_Type A_pos1": {0, 11}, "14_Type A_pos2": {0, 7, 11}, "14_Type A_pos3": {0, 5, 8, 11},
	"1_Type B_pos0": {0}, "1_Type B_pos1": {0}, "1_Type B_pos2": {0}, "1_Type B_pos3": {0},
	"2_Type B_pos0": {0}, "2_Type B_pos1": {0}, "2_Type B_pos2": {0}, "2_Type B_pos3": {0},
	"3_Type B_pos0": {0}, "3_Type B_pos1": {0}, "3_Type B_pos2": {0}, "3_Type B_pos3": {0},
	"4_Type B_pos0": {0}, "4_Type B_pos1": {0}, "4_Type B_pos2": {0}, "4_Type B_pos3": {0},
	"5_Type B_pos0": {0}, "5_Type B_pos1": {0, 4}, "5_Type B_pos2": {0, 4}, "5_Type B_pos3": {0, 4},
	"6_Type B_pos0": {0}, "6_Type B_pos1": {0, 4}, "6_Type B_pos2": {0, 4}, "6_Type B_pos3": {0, 4},
	"7_Type B_pos0": {0}, "7_Type B_pos1": {0, 4}, "7_Type B_pos2": {0, 4}, "7_Type B_pos3": {0, 4},
	"8_Type B_pos0": {0}, "8_Type B_pos1": {0, 6}, "8_Type B_pos2": {0, 3, 6}, "8_Type B_pos3": {0, 3, 6},
	"9_Type B_pos0": {0}, "9_Type B_pos1": {0, 6}, "9_Type B_pos2": {0, 3, 6}, "9_Type B_pos3": {0, 3, 6},
	"10_Type B_pos0": {0}, "10_Type B_pos1": {0, 8}, "10_Type B_pos2": {0, 4, 8}, "10_Type B_pos3": {0, 3, 6, 9},
	"11_Type B_pos0": {0}, "11_Type B_pos1": {0, 8}, "11_Type B_pos2": {0, 4, 8}, "11_Type B_pos3": {0, 3, 6, 9},
	"12_Type B_pos0": {0}, "12_Type B_pos1": {0, 10}, "12_Type B_pos2": {0, 5, 10}, "12_Type B_pos3": {0, 3, 6, 9},
	"13_Type B_pos0": {0}, "13_Type B_pos1": {0, 10}, "13_Type B_pos2": {0, 5, 10}, "13_Type B_pos3": {0, 3, 6, 9},
	"14_Type B_pos0": {0}, "14_Type B_pos1": {0, 10}, "14_Type B_pos2": {0, 5, 10}, "14_Type B_pos3": {0, 3, 6, 9},
}

// refer to 3GPP 38.211 vf30
//  Table 6.4.1.1.3-4: PUSCH DM-RS positions l- within a slot for double-symbol DM-RS and intra-slot frequency hopping disabled.
// key="ld_mapping type_additional position"
//  key = '%s_%s_%s' % (ld, self.nrDci01PuschTimeAllocMappingTypeComb.currentText(), self.nrDmrsDedPuschAddPosComb.currentText())
var DmrsPuschPosTwoSymbsWoIntraSlotFh = map[string][]int{
	"1_Type A_pos0": nil, "1_Type A_pos1": nil,
	"2_Type A_pos0": nil, "2_Type A_pos1": nil,
	"3_Type A_pos0": nil, "3_Type A_pos1": nil,
	"4_Type A_pos0": {0}, "4_Type A_pos1": {0},
	"5_Type A_pos0": {0}, "5_Type A_pos1": {0},
	"6_Type A_pos0": {0}, "6_Type A_pos1": {0},
	"7_Type A_pos0": {0}, "7_Type A_pos1": {0},
	"8_Type A_pos0": {0}, "8_Type A_pos1": {0},
	"9_Type A_pos0": {0}, "9_Type A_pos1": {0},
	"10_Type A_pos0": {0}, "10_Type A_pos1": {0, 8},
	"11_Type A_pos0": {0}, "11_Type A_pos1": {0, 8},
	"12_Type A_pos0": {0}, "12_Type A_pos1": {0, 8},
	"13_Type A_pos0": {0}, "13_Type A_pos1": {0, 10},
	"14_Type A_pos0": {0}, "14_Type A_pos1": {0, 10},
	"1_Type B_pos0": nil, "1_Type B_pos1": nil,
	"2_Type B_pos0": nil, "2_Type B_pos1": nil,
	"3_Type B_pos0": nil, "3_Type B_pos1": nil,
	"4_Type B_pos0": nil, "4_Type B_pos1": nil,
	"5_Type B_pos0": {0}, "5_Type B_pos1": {0},
	"6_Type B_pos0": {0}, "6_Type B_pos1": {0},
	"7_Type B_pos0": {0}, "7_Type B_pos1": {0},
	"8_Type B_pos0": {0}, "8_Type B_pos1": {0, 5},
	"9_Type B_pos0": {0}, "9_Type B_pos1": {0, 5},
	"10_Type B_pos0": {0}, "10_Type B_pos1": {0, 7},
	"11_Type B_pos0": {0}, "11_Type B_pos1": {0, 7},
	"12_Type B_pos0": {0}, "12_Type B_pos1": {0, 9},
	"13_Type B_pos0": {0}, "13_Type B_pos1": {0, 9},
	"14_Type B_pos0": {0}, "14_Type B_pos1": {0, 9},
}

// refer to 3GPP 38.211 vf30
//  Table 6.4.1.1.3-6: PUSCH DM-RS positions l- within a slot for single-symbol DM-RS and intra-slot frequency hopping enabled.
// key="ld per hop_mapping type_type a position_additional position_hop"
//  ld1 = math.floor(td / 2)
//  key1 = '%s_%s_%s_%s_1st' % (ld1, mappingType, self.nrMibDmRsTypeAPosComb.currentText()[3:] if mappingType == 'Type A' else '0', 'pos1' if self.nrDmrsDedPuschAddPosComb.currentText() != 'pos0' else 'pos0')
//  ld2 = td - math.floor(td / 2)
//  key2 = '%s_%s_%s_%s_2nd' % (ld2, mappingType, self.nrMibDmRsTypeAPosComb.currentText()[3:] if mappingType == 'Type A' else '0', 'pos1' if self.nrDmrsDedPuschAddPosComb.currentText() != 'pos0' else 'pos0')
var DmrsPuschPosOneSymbWithIntraSlotFh = map[string][]int{
	"1_Type A_2_pos0_1st": nil,
	"2_Type A_2_pos0_1st": nil,
	"3_Type A_2_pos0_1st": nil,
	"4_Type A_2_pos0_1st": {2},
	"5_Type A_2_pos0_1st": {2},
	"6_Type A_2_pos0_1st": {2},
	"7_Type A_2_pos0_1st": {2},
	"1_Type A_2_pos0_2nd": nil,
	"2_Type A_2_pos0_2nd": nil,
	"3_Type A_2_pos0_2nd": nil,
	"4_Type A_2_pos0_2nd": {0},
	"5_Type A_2_pos0_2nd": {0},
	"6_Type A_2_pos0_2nd": {0},
	"7_Type A_2_pos0_2nd": {0},
	"1_Type A_2_pos1_1st": nil,
	"2_Type A_2_pos1_1st": nil,
	"3_Type A_2_pos1_1st": nil,
	"4_Type A_2_pos1_1st": {2},
	"5_Type A_2_pos1_1st": {2},
	"6_Type A_2_pos1_1st": {2},
	"7_Type A_2_pos1_1st": {2, 6},
	"1_Type A_2_pos1_2nd": nil,
	"2_Type A_2_pos1_2nd": nil,
	"3_Type A_2_pos1_2nd": nil,
	"4_Type A_2_pos1_2nd": {0},
	"5_Type A_2_pos1_2nd": {0, 4},
	"6_Type A_2_pos1_2nd": {0, 4},
	"7_Type A_2_pos1_2nd": {0, 4},
	"1_Type A_3_pos0_1st": nil,
	"2_Type A_3_pos0_1st": nil,
	"3_Type A_3_pos0_1st": nil,
	"4_Type A_3_pos0_1st": {3},
	"5_Type A_3_pos0_1st": {3},
	"6_Type A_3_pos0_1st": {3},
	"7_Type A_3_pos0_1st": {3},
	"1_Type A_3_pos0_2nd": nil,
	"2_Type A_3_pos0_2nd": nil,
	"3_Type A_3_pos0_2nd": nil,
	"4_Type A_3_pos0_2nd": {0},
	"5_Type A_3_pos0_2nd": {0},
	"6_Type A_3_pos0_2nd": {0},
	"7_Type A_3_pos0_2nd": {0},
	"1_Type A_3_pos1_1st": nil,
	"2_Type A_3_pos1_1st": nil,
	"3_Type A_3_pos1_1st": nil,
	"4_Type A_3_pos1_1st": {3},
	"5_Type A_3_pos1_1st": {3},
	"6_Type A_3_pos1_1st": {3},
	"7_Type A_3_pos1_1st": {3},
	"1_Type A_3_pos1_2nd": nil,
	"2_Type A_3_pos1_2nd": nil,
	"3_Type A_3_pos1_2nd": nil,
	"4_Type A_3_pos1_2nd": {0},
	"5_Type A_3_pos1_2nd": {0, 4},
	"6_Type A_3_pos1_2nd": {0, 4},
	"7_Type A_3_pos1_2nd": {0, 4},
	"1_Type B_0_pos0_1st": {0},
	"2_Type B_0_pos0_1st": {0},
	"3_Type B_0_pos0_1st": {0},
	"4_Type B_0_pos0_1st": {0},
	"5_Type B_0_pos0_1st": {0},
	"6_Type B_0_pos0_1st": {0},
	"7_Type B_0_pos0_1st": {0},
	"1_Type B_0_pos0_2nd": {0},
	"2_Type B_0_pos0_2nd": {0},
	"3_Type B_0_pos0_2nd": {0},
	"4_Type B_0_pos0_2nd": {0},
	"5_Type B_0_pos0_2nd": {0},
	"6_Type B_0_pos0_2nd": {0},
	"7_Type B_0_pos0_2nd": {0},
	"1_Type B_0_pos1_1st": {0},
	"2_Type B_0_pos1_1st": {0},
	"3_Type B_0_pos1_1st": {0},
	"4_Type B_0_pos1_1st": {0},
	"5_Type B_0_pos1_1st": {0, 4},
	"6_Type B_0_pos1_1st": {0, 4},
	"7_Type B_0_pos1_1st": {0, 4},
	"1_Type B_0_pos1_2nd": {0},
	"2_Type B_0_pos1_2nd": {0},
	"3_Type B_0_pos1_2nd": {0},
	"4_Type B_0_pos1_2nd": {0},
	"5_Type B_0_pos1_2nd": {0, 4},
	"6_Type B_0_pos1_2nd": {0, 4},
	"7_Type B_0_pos1_2nd": {0, 4},
}

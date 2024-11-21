//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by conversion-gen. DO NOT EDIT.

package v2alpha1

import (
	url "net/url"
	unsafe "unsafe"

	datav0alpha1 "github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1"
	v0alpha1 "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1"
	dashboard "github.com/grafana/grafana/pkg/apis/dashboard"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AnnotationActions)(nil), (*dashboard.AnnotationActions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(a.(*AnnotationActions), b.(*dashboard.AnnotationActions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.AnnotationActions)(nil), (*AnnotationActions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(a.(*dashboard.AnnotationActions), b.(*AnnotationActions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AnnotationPermission)(nil), (*dashboard.AnnotationPermission)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_AnnotationPermission_To_dashboard_AnnotationPermission(a.(*AnnotationPermission), b.(*dashboard.AnnotationPermission), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.AnnotationPermission)(nil), (*AnnotationPermission)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_AnnotationPermission_To_v2alpha1_AnnotationPermission(a.(*dashboard.AnnotationPermission), b.(*AnnotationPermission), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Dashboard)(nil), (*dashboard.Dashboard)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_Dashboard_To_dashboard_Dashboard(a.(*Dashboard), b.(*dashboard.Dashboard), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.Dashboard)(nil), (*Dashboard)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_Dashboard_To_v2alpha1_Dashboard(a.(*dashboard.Dashboard), b.(*Dashboard), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DashboardAccess)(nil), (*dashboard.DashboardAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess(a.(*DashboardAccess), b.(*dashboard.DashboardAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.DashboardAccess)(nil), (*DashboardAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess(a.(*dashboard.DashboardAccess), b.(*DashboardAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DashboardList)(nil), (*dashboard.DashboardList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardList_To_dashboard_DashboardList(a.(*DashboardList), b.(*dashboard.DashboardList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.DashboardList)(nil), (*DashboardList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_DashboardList_To_v2alpha1_DashboardList(a.(*dashboard.DashboardList), b.(*DashboardList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DashboardVersionInfo)(nil), (*dashboard.DashboardVersionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardVersionInfo_To_dashboard_DashboardVersionInfo(a.(*DashboardVersionInfo), b.(*dashboard.DashboardVersionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.DashboardVersionInfo)(nil), (*DashboardVersionInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_DashboardVersionInfo_To_v2alpha1_DashboardVersionInfo(a.(*dashboard.DashboardVersionInfo), b.(*DashboardVersionInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DashboardVersionList)(nil), (*dashboard.DashboardVersionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardVersionList_To_dashboard_DashboardVersionList(a.(*DashboardVersionList), b.(*dashboard.DashboardVersionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.DashboardVersionList)(nil), (*DashboardVersionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_DashboardVersionList_To_v2alpha1_DashboardVersionList(a.(*dashboard.DashboardVersionList), b.(*DashboardVersionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DashboardWithAccessInfo)(nil), (*dashboard.DashboardWithAccessInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardWithAccessInfo_To_dashboard_DashboardWithAccessInfo(a.(*DashboardWithAccessInfo), b.(*dashboard.DashboardWithAccessInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.DashboardWithAccessInfo)(nil), (*DashboardWithAccessInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_DashboardWithAccessInfo_To_v2alpha1_DashboardWithAccessInfo(a.(*dashboard.DashboardWithAccessInfo), b.(*DashboardWithAccessInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LibraryPanel)(nil), (*dashboard.LibraryPanel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_LibraryPanel_To_dashboard_LibraryPanel(a.(*LibraryPanel), b.(*dashboard.LibraryPanel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.LibraryPanel)(nil), (*LibraryPanel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_LibraryPanel_To_v2alpha1_LibraryPanel(a.(*dashboard.LibraryPanel), b.(*LibraryPanel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LibraryPanelList)(nil), (*dashboard.LibraryPanelList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_LibraryPanelList_To_dashboard_LibraryPanelList(a.(*LibraryPanelList), b.(*dashboard.LibraryPanelList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.LibraryPanelList)(nil), (*LibraryPanelList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_LibraryPanelList_To_v2alpha1_LibraryPanelList(a.(*dashboard.LibraryPanelList), b.(*LibraryPanelList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LibraryPanelSpec)(nil), (*dashboard.LibraryPanelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec(a.(*LibraryPanelSpec), b.(*dashboard.LibraryPanelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.LibraryPanelSpec)(nil), (*LibraryPanelSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec(a.(*dashboard.LibraryPanelSpec), b.(*LibraryPanelSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LibraryPanelStatus)(nil), (*dashboard.LibraryPanelStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_LibraryPanelStatus_To_dashboard_LibraryPanelStatus(a.(*LibraryPanelStatus), b.(*dashboard.LibraryPanelStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.LibraryPanelStatus)(nil), (*LibraryPanelStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_LibraryPanelStatus_To_v2alpha1_LibraryPanelStatus(a.(*dashboard.LibraryPanelStatus), b.(*LibraryPanelStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VersionsQueryOptions)(nil), (*dashboard.VersionsQueryOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_VersionsQueryOptions_To_dashboard_VersionsQueryOptions(a.(*VersionsQueryOptions), b.(*dashboard.VersionsQueryOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*dashboard.VersionsQueryOptions)(nil), (*VersionsQueryOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_dashboard_VersionsQueryOptions_To_v2alpha1_VersionsQueryOptions(a.(*dashboard.VersionsQueryOptions), b.(*VersionsQueryOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*VersionsQueryOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_url_Values_To_v2alpha1_VersionsQueryOptions(a.(*url.Values), b.(*VersionsQueryOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v0alpha1.Unstructured)(nil), (*DashboardSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_Unstructured_To_v2alpha1_DashboardSpec(a.(*v0alpha1.Unstructured), b.(*DashboardSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*DashboardSpec)(nil), (*v0alpha1.Unstructured)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_DashboardSpec_To_v0alpha1_Unstructured(a.(*DashboardSpec), b.(*v0alpha1.Unstructured), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(in *AnnotationActions, out *dashboard.AnnotationActions, s conversion.Scope) error {
	out.CanAdd = in.CanAdd
	out.CanEdit = in.CanEdit
	out.CanDelete = in.CanDelete
	return nil
}

// Convert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions is an autogenerated conversion function.
func Convert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(in *AnnotationActions, out *dashboard.AnnotationActions, s conversion.Scope) error {
	return autoConvert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(in, out, s)
}

func autoConvert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(in *dashboard.AnnotationActions, out *AnnotationActions, s conversion.Scope) error {
	out.CanAdd = in.CanAdd
	out.CanEdit = in.CanEdit
	out.CanDelete = in.CanDelete
	return nil
}

// Convert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions is an autogenerated conversion function.
func Convert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(in *dashboard.AnnotationActions, out *AnnotationActions, s conversion.Scope) error {
	return autoConvert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(in, out, s)
}

func autoConvert_v2alpha1_AnnotationPermission_To_dashboard_AnnotationPermission(in *AnnotationPermission, out *dashboard.AnnotationPermission, s conversion.Scope) error {
	if err := Convert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(&in.Dashboard, &out.Dashboard, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_AnnotationActions_To_dashboard_AnnotationActions(&in.Organization, &out.Organization, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_AnnotationPermission_To_dashboard_AnnotationPermission is an autogenerated conversion function.
func Convert_v2alpha1_AnnotationPermission_To_dashboard_AnnotationPermission(in *AnnotationPermission, out *dashboard.AnnotationPermission, s conversion.Scope) error {
	return autoConvert_v2alpha1_AnnotationPermission_To_dashboard_AnnotationPermission(in, out, s)
}

func autoConvert_dashboard_AnnotationPermission_To_v2alpha1_AnnotationPermission(in *dashboard.AnnotationPermission, out *AnnotationPermission, s conversion.Scope) error {
	if err := Convert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(&in.Dashboard, &out.Dashboard, s); err != nil {
		return err
	}
	if err := Convert_dashboard_AnnotationActions_To_v2alpha1_AnnotationActions(&in.Organization, &out.Organization, s); err != nil {
		return err
	}
	return nil
}

// Convert_dashboard_AnnotationPermission_To_v2alpha1_AnnotationPermission is an autogenerated conversion function.
func Convert_dashboard_AnnotationPermission_To_v2alpha1_AnnotationPermission(in *dashboard.AnnotationPermission, out *AnnotationPermission, s conversion.Scope) error {
	return autoConvert_dashboard_AnnotationPermission_To_v2alpha1_AnnotationPermission(in, out, s)
}

func autoConvert_v2alpha1_Dashboard_To_dashboard_Dashboard(in *Dashboard, out *dashboard.Dashboard, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2alpha1_DashboardSpec_To_v0alpha1_Unstructured(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_Dashboard_To_dashboard_Dashboard is an autogenerated conversion function.
func Convert_v2alpha1_Dashboard_To_dashboard_Dashboard(in *Dashboard, out *dashboard.Dashboard, s conversion.Scope) error {
	return autoConvert_v2alpha1_Dashboard_To_dashboard_Dashboard(in, out, s)
}

func autoConvert_dashboard_Dashboard_To_v2alpha1_Dashboard(in *dashboard.Dashboard, out *Dashboard, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v0alpha1_Unstructured_To_v2alpha1_DashboardSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_dashboard_Dashboard_To_v2alpha1_Dashboard is an autogenerated conversion function.
func Convert_dashboard_Dashboard_To_v2alpha1_Dashboard(in *dashboard.Dashboard, out *Dashboard, s conversion.Scope) error {
	return autoConvert_dashboard_Dashboard_To_v2alpha1_Dashboard(in, out, s)
}

func autoConvert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess(in *DashboardAccess, out *dashboard.DashboardAccess, s conversion.Scope) error {
	out.Slug = in.Slug
	out.Url = in.Url
	out.CanSave = in.CanSave
	out.CanEdit = in.CanEdit
	out.CanAdmin = in.CanAdmin
	out.CanStar = in.CanStar
	out.CanDelete = in.CanDelete
	out.AnnotationsPermissions = (*dashboard.AnnotationPermission)(unsafe.Pointer(in.AnnotationsPermissions))
	return nil
}

// Convert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess is an autogenerated conversion function.
func Convert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess(in *DashboardAccess, out *dashboard.DashboardAccess, s conversion.Scope) error {
	return autoConvert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess(in, out, s)
}

func autoConvert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess(in *dashboard.DashboardAccess, out *DashboardAccess, s conversion.Scope) error {
	out.Slug = in.Slug
	out.Url = in.Url
	out.CanSave = in.CanSave
	out.CanEdit = in.CanEdit
	out.CanAdmin = in.CanAdmin
	out.CanStar = in.CanStar
	out.CanDelete = in.CanDelete
	out.AnnotationsPermissions = (*AnnotationPermission)(unsafe.Pointer(in.AnnotationsPermissions))
	return nil
}

// Convert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess is an autogenerated conversion function.
func Convert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess(in *dashboard.DashboardAccess, out *DashboardAccess, s conversion.Scope) error {
	return autoConvert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess(in, out, s)
}

func autoConvert_v2alpha1_DashboardList_To_dashboard_DashboardList(in *DashboardList, out *dashboard.DashboardList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]dashboard.Dashboard, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_Dashboard_To_dashboard_Dashboard(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v2alpha1_DashboardList_To_dashboard_DashboardList is an autogenerated conversion function.
func Convert_v2alpha1_DashboardList_To_dashboard_DashboardList(in *DashboardList, out *dashboard.DashboardList, s conversion.Scope) error {
	return autoConvert_v2alpha1_DashboardList_To_dashboard_DashboardList(in, out, s)
}

func autoConvert_dashboard_DashboardList_To_v2alpha1_DashboardList(in *dashboard.DashboardList, out *DashboardList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			if err := Convert_dashboard_Dashboard_To_v2alpha1_Dashboard(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_dashboard_DashboardList_To_v2alpha1_DashboardList is an autogenerated conversion function.
func Convert_dashboard_DashboardList_To_v2alpha1_DashboardList(in *dashboard.DashboardList, out *DashboardList, s conversion.Scope) error {
	return autoConvert_dashboard_DashboardList_To_v2alpha1_DashboardList(in, out, s)
}

func autoConvert_v2alpha1_DashboardVersionInfo_To_dashboard_DashboardVersionInfo(in *DashboardVersionInfo, out *dashboard.DashboardVersionInfo, s conversion.Scope) error {
	out.Version = in.Version
	out.ParentVersion = in.ParentVersion
	out.Created = in.Created
	out.CreatedBy = in.CreatedBy
	out.Message = in.Message
	return nil
}

// Convert_v2alpha1_DashboardVersionInfo_To_dashboard_DashboardVersionInfo is an autogenerated conversion function.
func Convert_v2alpha1_DashboardVersionInfo_To_dashboard_DashboardVersionInfo(in *DashboardVersionInfo, out *dashboard.DashboardVersionInfo, s conversion.Scope) error {
	return autoConvert_v2alpha1_DashboardVersionInfo_To_dashboard_DashboardVersionInfo(in, out, s)
}

func autoConvert_dashboard_DashboardVersionInfo_To_v2alpha1_DashboardVersionInfo(in *dashboard.DashboardVersionInfo, out *DashboardVersionInfo, s conversion.Scope) error {
	out.Version = in.Version
	out.ParentVersion = in.ParentVersion
	out.Created = in.Created
	out.CreatedBy = in.CreatedBy
	out.Message = in.Message
	return nil
}

// Convert_dashboard_DashboardVersionInfo_To_v2alpha1_DashboardVersionInfo is an autogenerated conversion function.
func Convert_dashboard_DashboardVersionInfo_To_v2alpha1_DashboardVersionInfo(in *dashboard.DashboardVersionInfo, out *DashboardVersionInfo, s conversion.Scope) error {
	return autoConvert_dashboard_DashboardVersionInfo_To_v2alpha1_DashboardVersionInfo(in, out, s)
}

func autoConvert_v2alpha1_DashboardVersionList_To_dashboard_DashboardVersionList(in *DashboardVersionList, out *dashboard.DashboardVersionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]dashboard.DashboardVersionInfo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2alpha1_DashboardVersionList_To_dashboard_DashboardVersionList is an autogenerated conversion function.
func Convert_v2alpha1_DashboardVersionList_To_dashboard_DashboardVersionList(in *DashboardVersionList, out *dashboard.DashboardVersionList, s conversion.Scope) error {
	return autoConvert_v2alpha1_DashboardVersionList_To_dashboard_DashboardVersionList(in, out, s)
}

func autoConvert_dashboard_DashboardVersionList_To_v2alpha1_DashboardVersionList(in *dashboard.DashboardVersionList, out *DashboardVersionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]DashboardVersionInfo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_dashboard_DashboardVersionList_To_v2alpha1_DashboardVersionList is an autogenerated conversion function.
func Convert_dashboard_DashboardVersionList_To_v2alpha1_DashboardVersionList(in *dashboard.DashboardVersionList, out *DashboardVersionList, s conversion.Scope) error {
	return autoConvert_dashboard_DashboardVersionList_To_v2alpha1_DashboardVersionList(in, out, s)
}

func autoConvert_v2alpha1_DashboardWithAccessInfo_To_dashboard_DashboardWithAccessInfo(in *DashboardWithAccessInfo, out *dashboard.DashboardWithAccessInfo, s conversion.Scope) error {
	if err := Convert_v2alpha1_Dashboard_To_dashboard_Dashboard(&in.Dashboard, &out.Dashboard, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_DashboardAccess_To_dashboard_DashboardAccess(&in.Access, &out.Access, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_DashboardWithAccessInfo_To_dashboard_DashboardWithAccessInfo is an autogenerated conversion function.
func Convert_v2alpha1_DashboardWithAccessInfo_To_dashboard_DashboardWithAccessInfo(in *DashboardWithAccessInfo, out *dashboard.DashboardWithAccessInfo, s conversion.Scope) error {
	return autoConvert_v2alpha1_DashboardWithAccessInfo_To_dashboard_DashboardWithAccessInfo(in, out, s)
}

func autoConvert_dashboard_DashboardWithAccessInfo_To_v2alpha1_DashboardWithAccessInfo(in *dashboard.DashboardWithAccessInfo, out *DashboardWithAccessInfo, s conversion.Scope) error {
	if err := Convert_dashboard_Dashboard_To_v2alpha1_Dashboard(&in.Dashboard, &out.Dashboard, s); err != nil {
		return err
	}
	if err := Convert_dashboard_DashboardAccess_To_v2alpha1_DashboardAccess(&in.Access, &out.Access, s); err != nil {
		return err
	}
	return nil
}

// Convert_dashboard_DashboardWithAccessInfo_To_v2alpha1_DashboardWithAccessInfo is an autogenerated conversion function.
func Convert_dashboard_DashboardWithAccessInfo_To_v2alpha1_DashboardWithAccessInfo(in *dashboard.DashboardWithAccessInfo, out *DashboardWithAccessInfo, s conversion.Scope) error {
	return autoConvert_dashboard_DashboardWithAccessInfo_To_v2alpha1_DashboardWithAccessInfo(in, out, s)
}

func autoConvert_v2alpha1_LibraryPanel_To_dashboard_LibraryPanel(in *LibraryPanel, out *dashboard.LibraryPanel, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*dashboard.LibraryPanelStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v2alpha1_LibraryPanel_To_dashboard_LibraryPanel is an autogenerated conversion function.
func Convert_v2alpha1_LibraryPanel_To_dashboard_LibraryPanel(in *LibraryPanel, out *dashboard.LibraryPanel, s conversion.Scope) error {
	return autoConvert_v2alpha1_LibraryPanel_To_dashboard_LibraryPanel(in, out, s)
}

func autoConvert_dashboard_LibraryPanel_To_v2alpha1_LibraryPanel(in *dashboard.LibraryPanel, out *LibraryPanel, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*LibraryPanelStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_dashboard_LibraryPanel_To_v2alpha1_LibraryPanel is an autogenerated conversion function.
func Convert_dashboard_LibraryPanel_To_v2alpha1_LibraryPanel(in *dashboard.LibraryPanel, out *LibraryPanel, s conversion.Scope) error {
	return autoConvert_dashboard_LibraryPanel_To_v2alpha1_LibraryPanel(in, out, s)
}

func autoConvert_v2alpha1_LibraryPanelList_To_dashboard_LibraryPanelList(in *LibraryPanelList, out *dashboard.LibraryPanelList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]dashboard.LibraryPanel)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2alpha1_LibraryPanelList_To_dashboard_LibraryPanelList is an autogenerated conversion function.
func Convert_v2alpha1_LibraryPanelList_To_dashboard_LibraryPanelList(in *LibraryPanelList, out *dashboard.LibraryPanelList, s conversion.Scope) error {
	return autoConvert_v2alpha1_LibraryPanelList_To_dashboard_LibraryPanelList(in, out, s)
}

func autoConvert_dashboard_LibraryPanelList_To_v2alpha1_LibraryPanelList(in *dashboard.LibraryPanelList, out *LibraryPanelList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]LibraryPanel)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_dashboard_LibraryPanelList_To_v2alpha1_LibraryPanelList is an autogenerated conversion function.
func Convert_dashboard_LibraryPanelList_To_v2alpha1_LibraryPanelList(in *dashboard.LibraryPanelList, out *LibraryPanelList, s conversion.Scope) error {
	return autoConvert_dashboard_LibraryPanelList_To_v2alpha1_LibraryPanelList(in, out, s)
}

func autoConvert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec(in *LibraryPanelSpec, out *dashboard.LibraryPanelSpec, s conversion.Scope) error {
	out.Type = in.Type
	out.PluginVersion = in.PluginVersion
	out.Title = in.Title
	out.Description = in.Description
	out.Options = in.Options
	out.FieldConfig = in.FieldConfig
	out.Datasource = (*datav0alpha1.DataSourceRef)(unsafe.Pointer(in.Datasource))
	out.Targets = *(*[]datav0alpha1.DataQuery)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec is an autogenerated conversion function.
func Convert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec(in *LibraryPanelSpec, out *dashboard.LibraryPanelSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_LibraryPanelSpec_To_dashboard_LibraryPanelSpec(in, out, s)
}

func autoConvert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec(in *dashboard.LibraryPanelSpec, out *LibraryPanelSpec, s conversion.Scope) error {
	out.Type = in.Type
	out.PluginVersion = in.PluginVersion
	out.Title = in.Title
	out.Description = in.Description
	out.Options = in.Options
	out.FieldConfig = in.FieldConfig
	out.Datasource = (*datav0alpha1.DataSourceRef)(unsafe.Pointer(in.Datasource))
	out.Targets = *(*[]datav0alpha1.DataQuery)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec is an autogenerated conversion function.
func Convert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec(in *dashboard.LibraryPanelSpec, out *LibraryPanelSpec, s conversion.Scope) error {
	return autoConvert_dashboard_LibraryPanelSpec_To_v2alpha1_LibraryPanelSpec(in, out, s)
}

func autoConvert_v2alpha1_LibraryPanelStatus_To_dashboard_LibraryPanelStatus(in *LibraryPanelStatus, out *dashboard.LibraryPanelStatus, s conversion.Scope) error {
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	out.Missing = in.Missing
	return nil
}

// Convert_v2alpha1_LibraryPanelStatus_To_dashboard_LibraryPanelStatus is an autogenerated conversion function.
func Convert_v2alpha1_LibraryPanelStatus_To_dashboard_LibraryPanelStatus(in *LibraryPanelStatus, out *dashboard.LibraryPanelStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_LibraryPanelStatus_To_dashboard_LibraryPanelStatus(in, out, s)
}

func autoConvert_dashboard_LibraryPanelStatus_To_v2alpha1_LibraryPanelStatus(in *dashboard.LibraryPanelStatus, out *LibraryPanelStatus, s conversion.Scope) error {
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	out.Missing = in.Missing
	return nil
}

// Convert_dashboard_LibraryPanelStatus_To_v2alpha1_LibraryPanelStatus is an autogenerated conversion function.
func Convert_dashboard_LibraryPanelStatus_To_v2alpha1_LibraryPanelStatus(in *dashboard.LibraryPanelStatus, out *LibraryPanelStatus, s conversion.Scope) error {
	return autoConvert_dashboard_LibraryPanelStatus_To_v2alpha1_LibraryPanelStatus(in, out, s)
}

func autoConvert_v2alpha1_VersionsQueryOptions_To_dashboard_VersionsQueryOptions(in *VersionsQueryOptions, out *dashboard.VersionsQueryOptions, s conversion.Scope) error {
	out.Path = in.Path
	out.Version = in.Version
	return nil
}

// Convert_v2alpha1_VersionsQueryOptions_To_dashboard_VersionsQueryOptions is an autogenerated conversion function.
func Convert_v2alpha1_VersionsQueryOptions_To_dashboard_VersionsQueryOptions(in *VersionsQueryOptions, out *dashboard.VersionsQueryOptions, s conversion.Scope) error {
	return autoConvert_v2alpha1_VersionsQueryOptions_To_dashboard_VersionsQueryOptions(in, out, s)
}

func autoConvert_dashboard_VersionsQueryOptions_To_v2alpha1_VersionsQueryOptions(in *dashboard.VersionsQueryOptions, out *VersionsQueryOptions, s conversion.Scope) error {
	out.Path = in.Path
	out.Version = in.Version
	return nil
}

// Convert_dashboard_VersionsQueryOptions_To_v2alpha1_VersionsQueryOptions is an autogenerated conversion function.
func Convert_dashboard_VersionsQueryOptions_To_v2alpha1_VersionsQueryOptions(in *dashboard.VersionsQueryOptions, out *VersionsQueryOptions, s conversion.Scope) error {
	return autoConvert_dashboard_VersionsQueryOptions_To_v2alpha1_VersionsQueryOptions(in, out, s)
}

func autoConvert_url_Values_To_v2alpha1_VersionsQueryOptions(in *url.Values, out *VersionsQueryOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	if values, ok := map[string][]string(*in)["version"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_int64(&values, &out.Version, s); err != nil {
			return err
		}
	} else {
		out.Version = 0
	}
	return nil
}

// Convert_url_Values_To_v2alpha1_VersionsQueryOptions is an autogenerated conversion function.
func Convert_url_Values_To_v2alpha1_VersionsQueryOptions(in *url.Values, out *VersionsQueryOptions, s conversion.Scope) error {
	return autoConvert_url_Values_To_v2alpha1_VersionsQueryOptions(in, out, s)
}

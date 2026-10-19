// SPDX-License-Identifier: AGPL-3.0-only
// Copyright (C) 2024-2026 Mikhail Kantur

package claims

const (
	UserId = "UserId"
	//UserName  = "UserName"
	AuthnType = "AuthnType" // { local | delegated }
	//AuthnProvider = "AuthProvider" // authentication provider name for multiple authenticators support
	TenantId = "TenantId"
)

const (
	AuthnTypeLocal     = "local"
	AuthnTypeDelegated = "delegated"
)

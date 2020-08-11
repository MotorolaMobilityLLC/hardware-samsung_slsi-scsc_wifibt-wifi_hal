package slsi_wifi_hal

import (
    "android/soong/android"
    "android/soong/cc"
)

func globalFlags(ctx android.BaseContext) []string {
    var cflags []string

    slsiwifihalconfig := ctx.DeviceConfig().WifiHalConfig()
    if (slsiwifihalconfig == "true") {
        cflags = append(cflags, "-DSLSI_WIFI_HAL_NL_ATTR_CONFIG")
    }
    return cflags
}

func deviceFlags(ctx android.BaseContext) []string {
    var cflags []string

    return cflags
}

func hostFlags(ctx android.BaseContext) []string {
    var cflags []string

    return cflags
}

func slsiWifiHalDefaults(ctx android.LoadHookContext) {
    type props struct {
        Target struct {
            Android struct {
                Cflags []string
                Enabled *bool
            }
            Host struct {
                Enabled *bool
            }
            Linux struct {
                Cflags []string
            }
            Darwin struct {
                Cflags []string
            }
        }
        Cflags []string
    }

    p := &props{}
    p.Cflags = globalFlags(ctx)
    p.Target.Android.Cflags = deviceFlags(ctx)
    h := hostFlags(ctx)
    p.Target.Linux.Cflags = h
    p.Target.Darwin.Cflags = h

    ctx.AppendProperties(p)
}

func init() {
    android.RegisterModuleType("slsi_wifi_hal_defaults", slsiWifiHalDefaultsFactory)
}

func slsiWifiHalDefaultsFactory() android.Module {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, slsiWifiHalDefaults)

    return module
}

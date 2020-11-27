package shared

import (
	"cockpitserver/simconnect"
)

var SimEventMap = map[string]simconnect.KeySimEvent{
	// Master
	"TOGGLE_AVIONICS_MASTER": simconnect.KeyToggleAvionicsMaster,
	"TOGGLE_MASTER_BATTERY":  simconnect.KeyToggleMasterBattery,
	// Toggle autopilot
	"AP_MASTER":           simconnect.KeyApMaster,
	"AUTOPILOT_ON":        simconnect.KeyAutopilotOn,
	"AUTOPILOT_OFF":       simconnect.KeyAutopilotOff,
	"SPOILERS_TOGGLE":     simconnect.KeySpoilersToggle,
	"SPOILERS_ARM_TOGGLE": simconnect.KeySpoilersArmToggle,
	// Autothrottle Heading Altitude Speed Approach modes
	"AUTO_THROTTLE_ARM":                simconnect.KeyAutoThrottleArm,
	"AP_HDG_HOLD":                      simconnect.KeyApPanelHeadingHold,
	"AP_ALT_HOLD":                      simconnect.KeyApAltHold,
	"AP_AIRSPEED_HOLD":                 simconnect.KeyApAirspeedHold,
	"AP_APR_HOLD":                      simconnect.KeyApAprHold,
	"HEADING_BUG_INC":                  simconnect.KeyHeadingBugInc,
	"HEADING_BUG_DEC":                  simconnect.KeyHeadingBugDec,
	"HEADING_BUG_SET":                  simconnect.KeyHeadingBugSet,
	"AP_VS_VAR_INC":                    simconnect.KeyApVsVarInc,
	"AP_VS_VAR_DEC":                    simconnect.KeyApVsVarDec,
	"AP_VS_VAR_SET_ENGLISH":            simconnect.KeyApVsVarSetEnglish,
	"AP_ALT_VAR_INC":                   simconnect.KeyApAltVarInc,
	"AP_ALT_VAR_DEC":                   simconnect.KeyApAltVarDec,
	"GEAR_TOGGLE":                      simconnect.KeyGearToggle,
	"TOGGLE_MASTER_BATTERY_ALTERNATOR": simconnect.KeyToggleMasterBattery,
	"MAGNETO":                          simconnect.KeyMagneto,
	"MINUS":                            simconnect.KeyMinus,
	"PLUS":                             simconnect.KeyPlus,
	// Toggle Lights
	"TOGGLE_BEACON_LIGHTS":      simconnect.KeyToggleBeaconLights,
	"STROBES_TOGGLE":            simconnect.KeyStrobesToggle,
	"TOGGLE_TAXI_LIGHTS":        simconnect.KeyToggleTaxiLights,
	"TOGGLE_LANDING_LIGHTS":     simconnect.KeyLandingLightsToggle,
	"TOGGLE_LOGO_LIGHTS":        simconnect.KeyToggleLogoLights,
	"TOGGLE_NAV_LIGHTS":         simconnect.KeyToggleNavLights,
	"TOGGLE_CABIN_LIGHTS":       simconnect.KeyToggleCabinLights,
	"TOGGLE_RECOGNITION_LIGHTS": simconnect.KeyToggleRecognitionLights,
	"PANEL_LIGHTS_TOGGLE":       simconnect.KeyPanelLightsToggle,
	"ALL_LIGHTS_TOGGLE":         simconnect.KeyAllLightsToggle,
}
var SimVarMap = map[string]simconnect.SimVar{
	// Master
	"ELECTRICAL MASTER BATTERY": simconnect.SimVarElectricalMasterBattery(),
	"AVIONICS MASTER SWITCH":    simconnect.SimVarAvionicsMasterSwitch(),

	// Navigation
	"PLANE ALTITUDE":                 simconnect.SimVarPlaneAltitude(simconnect.UnitFeet),
	"PLANE LATITUDE":                 simconnect.SimVarPlaneLatitude(simconnect.UnitDegrees), // you can force the units
	"PLANE LONGITUDE":                simconnect.SimVarPlaneLongitude(simconnect.UnitDegrees),
	"INDICATED ALTITUDE":             simconnect.SimVarIndicatedAltitude(simconnect.UnitFeet),
	"AIRSPEED INDICATED":             simconnect.SimVarAirspeedIndicated(),
	"VERTICAL SPEED":                 simconnect.SimVarVerticalSpeed(simconnect.UnitFeet),
	"HEADING INDICATOR":              simconnect.SimVarHeadingIndicator(simconnect.UnitDegrees),
	"PLANE HEADING DEGREES TRUE":     simconnect.SimVarPlaneHeadingDegreesTrue(simconnect.UnitDegrees),
	"PLANE HEADING DEGREES MAGNETIC": simconnect.SimVarPlaneHeadingDegreesMagnetic(simconnect.UnitDegrees),
	"DELTA HEADING RATE":             simconnect.SimVarDeltaHeadingRate(simconnect.UnitDegrees),
	"SPOILERS HANDLE POSITION":       simconnect.SimVarSpoilersHandlePosition(),
	// Lights
	"LIGHT BEACON":  simconnect.SimVarLightBeacon(),
	"LIGHT TAXI":    simconnect.SimVarLightTaxi(),
	"LIGHT WING":    simconnect.SimVarLightWing(),
	"LIGHT LOGO":    simconnect.SimVarLightLogo(),
	"LIGHT PANEL":   simconnect.SimVarLightPanel(),
	"LIGHT STROBE":  simconnect.SimVarLightStrobe(),
	"LIGHT LANDING": simconnect.SimVarLightLanding(),
	"LIGHT NAV":     simconnect.SimVarLightNav(),
	"LIGHT CABIN":   simconnect.SimVarLightCabin(),
	// Misc
	"GEAR POSITION":       simconnect.SimVarGearPosition(),
	"SIM ON GROUND":       simconnect.SimVarSimOnGround(),
	"PLANE PITCH DEGREES": simconnect.SimVarPlanePitchDegrees(),
	"PLANE BANK DEGREES":  simconnect.SimVarPlaneBankDegrees(),
	// Autopilot
	"AUTOPILOT AVAILABLE":         simconnect.SimVarAutopilotAvailable(),
	"AUTOPILOT MASTER":            simconnect.SimVarAutopilotMaster(),
	"AUTOPILOT THROTTLE ARM":      simconnect.SimVarAutopilotThrottleArm(),
	"AUTOPILOT HEADING LOCK":      simconnect.SimVarAutopilotHeadingLock(),
	"AUTOPILOT HEADING LOCK DIR":  simconnect.SimVarAutopilotHeadingLockDir(simconnect.UnitDegrees),
	"AUTOPILOT AIRSPEED HOLD":     simconnect.SimVarAutopilotAirspeedHold(),
	"AUTOPILOT AIRSPEED HOLD VAR": simconnect.SimVarAutopilotAirspeedHoldVar(),
	"AUTOPILOT APPROACH HOLD":     simconnect.SimVarAutopilotApproachHold(),
	"AUTOPILOT ALTITUDE LOCK":     simconnect.SimVarAutopilotAltitudeLock(),
	"AUTOPILOT ALTITUDE LOCK VAR": simconnect.SimVarAutopilotAltitudeLockVar(),
	// Radio Comms
	"COM ACTIVE FREQUENCY":  simconnect.SimVarComActiveFrequency(),
	"NAV ACTIVE FREQUENCY":  simconnect.SimVarNavActiveFrequency(),
	"COM STANDBY FREQUENCY": simconnect.SimVarComStandbyFrequency(),
	"NAV STANDBY FREQUENCY": simconnect.SimVarNavStandbyFrequency(),
}

type SimVariable map[string]float64

func CreateSimConnectConnection(appName string) (*simconnect.EasySimConnect, <-chan SimVariable) {

	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	sc.SetLoggerLevel(simconnect.LogInfo)
	c, err := sc.Connect(appName)
	if err != nil {
		panic(err)
	}
	<-c // wait connection confirmation
	for {
		if <-sc.ConnectSysEventSim() {
			break // wait sim start
		}
	}

	// Collect all registered sim variables
	var simvars []simconnect.SimVar = []simconnect.SimVar{}
	for _, variable := range SimVarMap {
		simvars = append(simvars, variable)
	}
	// Connect to all variables
	cSimVar, err := sc.ConnectToSimVar(
		simvars...,
	)

	// Structure the received sim variable to pass to a live stream
	simVariableChannel := make(chan SimVariable)
	go func() {
		for {
			result := <-cSimVar
			t := SimVariable{}
			for _, simVar := range result {
				f, err := simVar.GetFloat64()
				if err != nil {
					panic(err)
				}
				t[simVar.Name] = f
			}
			simVariableChannel <- t
		}
	}()
	return sc, simVariableChannel
}

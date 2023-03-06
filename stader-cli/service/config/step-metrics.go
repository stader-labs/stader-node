package config

import cfgtypes "github.com/stader-labs/stader-node/shared/types/config"

func createMetricsStep(wiz *wizard, currentStep int, totalSteps int) *choiceWizardStep {

	helperText := "Would you like to enable the Stader Node's metrics monitoring system? This will monitor things such as hardware stats (CPU usage, RAM usage, free disk space), your validator stats, stats about your node such as ETH rewards, and much more.\n\nNone of this information will be sent to any remote servers for collection an analysis; this is purely for your own usage on your node."

	show := func(modal *choiceModalLayout) {
		wiz.md.setPage(modal.page)
		if wiz.md.Config.EnableMetrics.Value == false {
			modal.focus(0)
		} else {
			modal.focus(1)
		}
	}

	done := func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 1 {
			wiz.md.Config.EnableMetrics.Value = true
		} else {
			wiz.md.Config.EnableMetrics.Value = false
		}
		if wiz.md.Config.Smartnode.Network.Value.(cfgtypes.Network) == cfgtypes.Network_Zhejiang {
			wiz.finishedModal.show()
		} else {
			wiz.mevModeModal.show()
		}
	}

	back := func() {
		wiz.useFallbackModal.show()
	}

	return newChoiceStep(
		wiz,
		currentStep,
		totalSteps,
		helperText,
		[]string{"No", "Yes"},
		[]string{},
		76,
		"Metrics",
		DirectionalModalHorizontal,
		show,
		done,
		back,
		"step-metrics",
	)

}

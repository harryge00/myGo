// podKiller launches a goroutine to kill a pod received from the channel if
// another goroutine isn't already in action.
func (kl *Kubelet) podKiller() {
	killing := sets.NewString()
	// guard for the killing set
	lock := sync.Mutex{}
	for {
		select {
		case podPair, ok := <-kl.podKillingCh:
			if !ok {
				return
			}

			runningPod := podPair.RunningPod
			apiPod := podPair.APIPod

			lock.Lock()
			exists := killing.Has(string(runningPod.ID))
			if !exists {
				killing.Insert(string(runningPod.ID))
			}
			lock.Unlock()

			if !exists {
				go func(apiPod *v1.Pod, runningPod *kubecontainer.Pod) {
					glog.V(2).Infof("Killing unwanted pod %q", runningPod.Name)
					err := kl.killPod(apiPod, runningPod, nil, nil)
					if err != nil {
						glog.Errorf("Failed killing the pod %q: %v", runningPod.Name, err)
					}
					lock.Lock()
					killing.Delete(string(runningPod.ID))
					lock.Unlock()
				}(apiPod, runningPod)
			}
		}
	}
}
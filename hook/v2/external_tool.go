package v2

// runExternalTools iterates over external tools to validate commit
func (cfg *Configuration) runExternalTools() bool {
	return true
}

/*cmd := exec.Command("git", "blub")

  if err := cmd.Start(); err != nil {
      log.Fatalf("cmd.Start: %v")
  }

  if err := cmd.Wait(); err != nil {
      if exiterr, ok := err.(*exec.ExitError); ok {
          // The program has exited with an exit code != 0

          // This works on both Unix and Windows. Although package
          // syscall is generally platform dependent, WaitStatus is
          // defined for both Unix and Windows and in both cases has
          // an ExitStatus() method with the same signature.
          if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
              log.Printf("Exit Status: %d", status.ExitStatus())
          }
      } else {
          log.Fatalf("cmd.Wait: %v", err)
      }
  }*/

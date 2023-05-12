#include <signal.h>

int main() {
  // Create a bit vector to represent the set of enabled signals.
  int enabled_signals = 0;

  // Enable the SIGINT signal.
  signal(SIGINT, handle_sigint);
  enabled_signals |= (1 << SIGINT);

  // Enable the SIGTSTP signal.
  signal(SIGTSTP, handle_sigstop);
  enabled_signals |= (1 << SIGTSTP);

  // Loop forever, handling any enabled signals.
  while (1) {
    // Wait for a signal.
    pause();

    // Check the signal that was received.
    switch (sig) {
      case SIGINT:
        // Handle the SIGINT signal.
        handle_sigint();
        break;

      case SIGTSTP:
        // Handle the SIGTSTP signal.
        handle_sigstop();
        break;

      default:
        // Ignore any other signals.
        break;
    }
  }

  return 0;
}

void handle_sigint() {
  // Do something when the SIGINT signal is received.
}

void handle_sigstop() {
  // Do something when the SIGTSTP signal is received.
}

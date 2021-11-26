#include "hackrf.h"
#include "_cgo_export.h"
#include "extras.h"

int rx_callback_helper(hackrf_transfer* transfer) {
    return rxCallbackHelper(transfer);
}

int tx_callback_helper(hackrf_transfer* transfer) {
    return txCallbackHelper(transfer);
}

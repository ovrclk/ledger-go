/*******************************************************************************
*   (c) 2018 ZondaX GmbH
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
********************************************************************************/

package ledger_go

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mux sync.Mutex

func Test_CountLedgerDevices(t *testing.T) {
	mux.Lock()
	defer mux.Unlock()

	ledgerAdmin := NewLedgerAdmin()
	count := ledgerAdmin.CountDevices()
	assert.True(t, count > 0)
}

func Test_ListDevices(t *testing.T) {
	mux.Lock()
	defer mux.Unlock()

	ledgerAdmin := NewLedgerAdmin()
	_, err := ledgerAdmin.ListDevices()
	require.NoError(t, err)
}

func Test_GetLedger(t *testing.T) {
	mux.Lock()
	defer mux.Unlock()

	ledgerAdmin := NewLedgerAdmin()
	count := ledgerAdmin.CountDevices()
	require.True(t, count > 0)

	ledger, err := ledgerAdmin.Connect(0)
	assert.NoError(t, err)
	assert.NotNil(t, ledger)

	defer func() {
		_ = ledger.Close()
	}()
}

func Test_BasicExchange(t *testing.T) {
	mux.Lock()
	defer mux.Unlock()

	ledgerAdmin := NewLedgerAdmin()
	count := ledgerAdmin.CountDevices()
	require.True(t, count > 0)

	ledger, err := ledgerAdmin.Connect(0)
	assert.NoError(t, err)
	assert.NotNil(t, ledger)

	defer func() {
		_ = ledger.Close()
	}()

	// Call device info (this should work in main menu and many apps)
	message := []byte{0xE0, 0x01, 0, 0, 0}

	for i := 0; i < 10; i++ {
		response, err := ledger.Exchange(message)

		if err != nil {
			fmt.Printf("iteration %d\n", i)
			t.Fatalf("Error: %s", err.Error())
		}

		require.True(t, len(response) > 0)
	}
}

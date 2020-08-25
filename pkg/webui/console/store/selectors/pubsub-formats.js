// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { GET_PUBSUB_FORMATS_BASE } from '@console/store/actions/pubsub-formats'

import { createFetchingSelector } from './fetching'
import { createErrorSelector } from './error'

const selectPubsubFormatsStore = state => state.pubsubFormats

export const selectPubsubFormats = function(state) {
  const store = selectPubsubFormatsStore(state)

  return store.formats || {}
}

export const selectPubsubFormatsError = createErrorSelector(GET_PUBSUB_FORMATS_BASE)
export const selectPubsubFormatsFetching = createFetchingSelector(GET_PUBSUB_FORMATS_BASE)
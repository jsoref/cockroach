// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package opgen

import (
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scpb"
)

func init() {
	opRegistry.register(
		(*scpb.SequenceOwnedBy)(nil),
		scpb.Target_DROP,
		scpb.Status_PUBLIC,
		to(scpb.Status_ABSENT,
			// TODO(ajwerner): This probably should be marked as non-revertible.
			minPhase(scop.PreCommitPhase),
			emit(func(this *scpb.SequenceOwnedBy) scop.Op {
				return &scop.RemoveSequenceOwnedBy{
					TableID: this.SequenceID,
				}
			})),
	)
}

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers_test

import (
	"encoding/base64"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/alex123012/database-users-operator/api/v1alpha1"
)

const (
	databaseBindingCreationTimeout = 3 * time.Second
	// databaseBindingDeletionTimeout = 3 * time.Second.
)

var _ = Describe("DatabaseBindingController", func() {
	var (
		user            *v1alpha1.User
		secret          *v1.Secret
		database        *v1alpha1.Database
		databaseBinding *v1alpha1.DatabaseBinding
		namespace       = "default"
	)
	Context("default behaviour", func() {
		BeforeEach(func() {
			user, secret, database, databaseBinding = databaseBindingBundle(namespace)
			Expect(k8sClient.Create(ctx, user)).To(Succeed())
			Expect(k8sClient.Create(ctx, secret)).To(Succeed())
			Expect(k8sClient.Create(ctx, database)).To(Succeed())
			Expect(k8sClient.Create(ctx, databaseBinding)).To(Succeed())
			waitForDatabaseBindingReady(databaseBinding)
		})

		AfterEach(func() {
			for _, o := range []client.Object{databaseBinding, user, database} {
				Expect(k8sClient.Delete(ctx, o)).To(Succeed())
				Eventually(func(o client.Object) bool {
					err := k8sClient.Get(ctx, types.NamespacedName{Name: o.GetName(), Namespace: o.GetNamespace()}, o)
					return apierrors.IsNotFound(err)
				}, 5).WithArguments(o).Should(BeTrue())
			}
			fakeDBCreatorDB.ResetDB()
		})

		It("works", func() {
			queries := fakeDBCreatorDB.Conn.Queries()
			By("Created user in DB", func() {
				Expect(queries[`CREATE USER "user-1" WITH PASSWORD 'mysupersecretpass'`]).NotTo(Equal(0))
			})
		})
	})
})

func waitForDatabaseBindingReady(databaseBinding *v1alpha1.DatabaseBinding) {
	EventuallyWithOffset(1, func() string {
		databaseBindingCreated := v1alpha1.DatabaseBinding{}
		if err := k8sClient.Get(
			ctx,
			types.NamespacedName{Name: databaseBinding.Name, Namespace: databaseBinding.Namespace},
			&databaseBindingCreated,
		); err != nil {
			return fmt.Sprintf("%v+", err)
		}

		if !databaseBindingCreated.Status.Summary.Ready {
			return "not ready"
		}

		return "ready"
	}, databaseBindingCreationTimeout, 1*time.Second).Should(Equal("ready"))
}

func databaseBindingBundle(namespace string) (*v1alpha1.User, *v1.Secret, *v1alpha1.Database, *v1alpha1.DatabaseBinding) {
	database := &v1alpha1.Database{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "database-1",
			Namespace: namespace,
		},
		Spec: v1alpha1.DatabaseSpec{
			Type: v1alpha1.PostgreSQL,
		},
	}

	user := &v1alpha1.User{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "user-1",
			Namespace: namespace,
		},
		PasswordSecret: v1alpha1.Secret{
			Secret: v1alpha1.NamespacedName{
				Name:      "user-password-1",
				Namespace: namespace,
			},
			Key: "pass",
		},
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "user-password-1",
			Namespace: namespace,
		},
		Data: map[string][]byte{
			"pass": []byte(base64.StdEncoding.EncodeToString([]byte("mysupersecretpass"))),
		},
	}

	databaseBinding := &v1alpha1.DatabaseBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "database-binding",
			Namespace: namespace,
		},
		Spec: v1alpha1.DatabaseBindingSpec{
			Database: v1alpha1.NamespacedName{
				Namespace: namespace,
				Name:      "database-1",
			},
			User: v1alpha1.NamespacedName{
				Namespace: namespace,
				Name:      "user-1",
			},
		},
	}

	return user, secret, database, databaseBinding
}

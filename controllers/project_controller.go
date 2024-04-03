/*
Copyright 2024.

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

package controllers

import (
	"context"
	"fmt"
	"strconv"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	manageablesv1alpha1 "github.com/dominik-robert/manager/api/v1alpha1"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=manageables.dorosoft.de,resources=projects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=manageables.dorosoft.de,resources=projects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=manageables.dorosoft.de,resources=projects/finalizers,verbs=update

//+kubebuilder:rbac:groups=manageables.dorosoft.de,resources=people,verbs=get;list;update;patch
//+kubebuilder:rbac:groups=manageables.dorosoft.de,resources=people/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Project object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ProjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	projects := &manageablesv1alpha1.ProjectList{}
	people := &manageablesv1alpha1.PeopleList{}

	opts := []client.ListOption{
		client.InNamespace(req.NamespacedName.Namespace),
	}
	err := r.List(ctx, projects, opts...)
	if err != nil {
		fmt.Println(err)
	}
	err = r.List(ctx, people, opts...)
	if err != nil {
		fmt.Println(err)
	} else {
		for i, project := range projects.Items {
			matchingPeople := make([]manageablesv1alpha1.MatchingPeople, 0)
			for _, person := range people.Items {
				skills, perc, perWAlt, percWExp := lookupSameSkills(project.Spec.NeededSkills, person.Spec.Skills)
				if true || len(skills) != 0 || perc != "0.00" {
					matchedPeopleDetails := manageablesv1alpha1.MatchingPeople{
						PercentageStrict:              perc,
						PercentageWithoutAlternatives: perWAlt,
						PercentageWithoutExperience:   percWExp,
						PersonName:                    person.Name,
						Skills:                        skills,
					}
					matchingPeople = append(matchingPeople, matchedPeopleDetails)
				}
			}
			projects.Items[i].Status.PossiblePeople = matchingPeople
			r.Status().Update(ctx, &projects.Items[i])
		}
	}
	return ctrl.Result{}, nil
}

func lookupSameSkills(projectSkills, personSkills []manageablesv1alpha1.SkillSpecDetails) ([]manageablesv1alpha1.SkillResolved, string, string, string) {
	tests := make(map[string]*manageablesv1alpha1.SkillResolved)

	// convert list of skill projects to a map for faster and easier access
	for _, value := range projectSkills {
		tests[value.Name] = &manageablesv1alpha1.SkillResolved{
			SkillSpecs: value,
		}
	}

	// initialize counting variables for all matching skills to calculate the percentage
	matchesStrict := 0
	matchesAlt := 0
	matchesWithoutExp := 0

	fmt.Println("MAP", tests)
	// iterate through every person for skill compare
	for _, value := range personSkills {
		item, exists := tests[value.Name]
		fmt.Println("Value Name", value.Name)
		// check if item exists in the map
		if exists {
			// people has the skill which is in the project
			// calculate if enough experience is there
			timeExpected := parseTimeToDays(item.SkillSpecs.Experience)
			timeAvailable := parseTimeToDays(value.Experience)
			item.SkillAvailable = true
			if timeAvailable >= timeExpected {
				item.ExperienceFulfilled = true
				matchesStrict++
			}
			matchesWithoutExp++
		} else {
			// item does not exists so skill is not known. Lookup for the alternatives
			// TODO(implement)
		}
	}
	percentage := fmt.Sprintf("%.2f", float32(matchesWithoutExp)/float32(len(projectSkills))*100)
	percentageWithoutAlternatives := fmt.Sprintf("%.2f", float32(matchesAlt)/float32(len(projectSkills))*100)
	percentageWithoutExperience := fmt.Sprintf("%.2f", float32(matchesStrict)/float32(len(projectSkills))*100)

	// convert map to slice
	skillSet := make([]manageablesv1alpha1.SkillResolved, len(tests))
	iter := 0
	for _, value := range tests {
		skillSet[iter] = *value
		iter++
	}

	return skillSet, percentage, percentageWithoutAlternatives, percentageWithoutExperience
}

func parseTimeToDays(timeSpec string) int {

	if len(timeSpec) == 0 || timeSpec == "" {
		return 0
	}

	suffix := timeSpec[len(timeSpec)-1:]
	number, _ := strconv.Atoi(timeSpec[:len(timeSpec)-1])
	switch suffix {
	case "y":
		return number * 365
	case "m":
		return number * 30
	default:
		return 0
	}

}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&manageablesv1alpha1.Project{}).
		Complete(r)
}

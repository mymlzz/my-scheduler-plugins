package plugins

import(
  "fmt"
  v1 "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/runtime"
  "k8s.io/klog/v2"
  "k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "myscheduler-plugin"

type Args struct {
  FavoriteColor string `json:"favorite_color,omitempty"`
  FavoriteNumber int `json:"favorite_number,omitempty"`
  ThanksTo string `json:"thanks_to,omitempty"`
}

type Myscheduler struct {
  args *Args
  handle framework.FrameworkHandle
}

func (s *Myscheduler) Name() string {
  return Name
}

func (s *Myscheduler) PreFilter(pc *framework.PluginContext, pod *v1.Pod) *framework.Status{
  klog.Infof("prefilter pod: %v", pod.Name)
  return framework.NewStatus(framework.Success, "")
}

func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error){
  args := &Args{}
  if err := framework.DecodeInto(configuration, args); err != nil{
    return nil, err
  }
  klog.infof("get plugin config args: %+v",args)
  return &Sample{
    args: args,
    handle: f,
  }
}

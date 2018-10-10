job "my-service" {
  datacenters = ["dc1"]
  type        = "service"

  #  A group defines a series of tasks that should be co-located on the same client (host)
  group "my-service" {
    count = 2

    # Task is a unit of work and is executed by a driver
    task "my-service" {
      driver = "raw_exec"

      config {
        command = "my-service"
      }

      # Configuration is specific to each driver
      artifact {
        source = "http://127.0.0.1:8000/my-service"
      }

      # Specify the maximum resources required to run the task
      resources {
        cpu    = 200
        memory = 100

        network {
          port "http" {}
        }
      }

      # Specifies integration with Consul
      service {
        port = "http"
        name = "my-service"

        tags = [
          "dev",
        ]

        check {
          type     = "http"
          port     = "http"
          path     = "/hello"
          interval = "45s"
          timeout  = "30s"
        }
      }
    }
  }
}

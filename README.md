# Dahua Panel 

[![Build Status](https://travis-ci.com/denisrudov/dahua_panel.svg?branch=master)](https://travis-ci.com/denisrudov/dahua_panel)

Simple client to manage **VTO2111D** Dahua Panel

## Getting Started

```go get github.com/denisrudov/dahua_panel```

or even simpler

```import github.com/denisrudov/dahua_panel ``` 

in your code and then

``` dep ensure ```
 

## Authors

* **Denis Rudov** 

feel free to contribute :)


### How to use

Just create client with credentials and IP address of a dahua panel

``` client := NewDahuaClient("admin", "adminpassword", "192.168.0.91")```


To log in

``` client.Login() ```

To update maintain parameters

```
if client.Login() {
    maintainParams := dahua_panel.NewMaintainParams()

    if err := client.UpdateMaintainParams(maintainParams); err != nil {
       log.Println(err)
    }
}
```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

# eddy the systemd unit file generator

## Create your service file

For example make some command a (user mode) service
 ```bash
 eddy create service  \ 
    --description "my sleeping daemon" \
    --exec-start "sleep 5m" \
    --install-required-by "multi-user.target" \
    > ~/.config/systemd/user/mysleeping.service

 ```

Now enable and run the service:
```bash
systemctl --user enable mysleeping.service
systemctl --user start mysleeping.service

```

Get eddy's bash completion
```bash
source <(eddy completion) 
```

## Build and install it
```bash
make build
chmod +x eddy
# copy to your favorite PATH location
sudo cp eddy /usr/local/bin/
```

## Example usage: run a postgresql 10 container as system service using podman
```bash
eddy create service \
    --description "my pg10 dev service" \
    --exec-start "podman run -it postgres:10.6" \
    --install-required-by "multi-user.target" \
    > ~/.config/systemd/user/pg10.service
    
systemctl --user enable pg10.service
systemctl --user start pg10.service

# stop when uneeded, it just works
systemctl --user stop pg10.service

```
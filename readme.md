# Web of Things Framework for Go

This will be an experimental implementation of the Web of Things Framework written in the Go programming language. This project is a sister of the [NodeJS server for the Web of Things](https://github.com/w3c/web-of-things-framework).

A hundred billion IoT devices are expected to be deployed over the next ten years. There are many IoT platforms, and an increasing number of IoT technologies. However, the IoT products are currently beset by silos and a lack of interoperability, which blocks the benefits of the network effect from taking hold.  At W3C, we are exploring the potential for bridging IoT platforms and devices through the World Wide Web via a new class of Web servers that are connected through a common framework, and available in a variety of scales from microcontrollers, to smart phones and home/office hubs, and cloud-based server farms.

This framework involves virtual objects ("things") as proxies for physical and abstract entities. These things are modelled in terms of metadata, events, properties and actions, with bindings to scripting APIs and a variety of protocols, since no one protocol will fulfil all needs. This server will eventually support bindings to HTTP, WebSockets, CoAP, MQTT, XMPP and AMQP.

## Technical Recap.

The Web of Things starts with URIs for “thing” descriptions as Linked Data expressed in JSON-LD. Things are modelled in terms of their events, properties and actions. Scripts use this URI to ask the server to create a local proxy object in the script’s execution space for the designated thing. The server takes care of the protocol details, and can use which protocol best suits its needs, e.g. HTTP, WebSockets, CoAP, MQTT, XMPP or AMQP.  The thing metadata and the target server metadata enable the server hosting the proxy to figure out which protocol to use for the messaging between the proxy and the thing it proxies.

You can have chains of proxies across different servers, which can range from microcontrollers to cloud-based server farms. Whilst REST based messages are used, the URI paths are really a matter for each server (just as in HTTP).  Discovery involves a range of techniques: mDNS and UPnP on local networks, Bluetooth and BLE beacons, NFC, barcodes, cloud based device registries, social networks of people and things, and by following links between things from the dependencies stated in their descriptions. Servers may also expose a list of the things they host. Privacy will be based upon access control, and terms & conditions that enable data owners to control who can access their data and for what purpose.

The [Go programming language](https://golang.org) is a good choice for high performance servers on powerful platforms, and offers a very clean approach to concurrent programming.

## Prerequisites

 1. Git
 2. Go
 3. *Go modules for the protocols*
 
## Installation

The starting point is to install Git, see:

  http://git-scm.com/book/en/v2/Getting-Started-Installing-Git

Next create a copy of this directory and change to it as follows:

```
  git clone https://github.com/w3c/wot-go
  cd wot-go
```

## Contributing

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/w3c/web-of-things-framework?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

We welcome contributions. If you find a bug in the source code or a mistake in the documentation, you can help us by submitting an issue to our [GitHub repository](https://github.com/w3c/thoosa), and likewise if you have suggestions for new features. Even better you can submit a Pull Request with a fix. We also have a Gitter chat channel, see above link.

We encourage you to join the W3C [Web of Things Community Group](https://www.w3.org/community/wot/) where contribution and discussions happen. Anyone can join and there are no fees.

The amount of time you contribute and the areas in which you contribute is up to you. 

## License

(The MIT License)

Copyright (c) 2015 Dave Raggett &lt;dsr@w3.org&gt;

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

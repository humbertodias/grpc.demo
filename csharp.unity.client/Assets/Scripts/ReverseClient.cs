using Grpc.Core;
using Grpc.Net.Client.Web;
using Grpc.Net.Client;
using System.Net.Http;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using System;
using Reverse;

public class ReverseClient : MonoBehaviour
{
    
    public Button Button;

    public TMP_Text Text;

    public string host = "http://localhost:50051";

    private GrpcChannel channel;
    private ReverseService.ReverseServiceClient client;

    void Start()
    {
        Button.onClick.AddListener(OnHello);
        var options = new GrpcChannelOptions();
        var handler = new GrpcWebHandler(GrpcWebMode.GrpcWeb, new HttpClientHandler());
        options.HttpHandler = handler;
        options.Credentials = ChannelCredentials.Insecure;
        channel = GrpcChannel.ForAddress(host, options);
        client = new ReverseService.ReverseServiceClient(channel);
    }
    private void OnDestroy()
    {
        channel.Dispose();
    }

    public void OnHello()
    {
        var reply = client.ReverseString(new ReverseRequest
        {
            Data = "Hello, World from Unity"
        });
        Debug.Log(reply.ToString());
        Text.text = $"[{DateTime.Now}] {reply}";
    }
}

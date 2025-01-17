using AuthService;
using Azure.Core;
using Grpc.Net.Client;
using Microsoft.Data.SqlClient;
using System.Data;
using System.Runtime.Serialization;
using System.ServiceModel;
using Microsoft.AspNetCore.Http;

namespace CreateInvoice.BusinessLogic
{
    [DataContract]
    public class Invoice
    {
        [DataMember]
        public DateTime BillingDate { get; set; }

        [DataMember]
        public string PaymentMethod { get; set; } = string.Empty;

        [DataMember]
        public decimal Amount { get; set; }

        [DataMember]
        public decimal ShippingFee { get; set; }

        [DataMember]
        public int ClientId { get; set; }

        [DataMember]
        public decimal Tax { get; set; }
    }


    [ServiceContract]
    public interface ISoapService
    {
        [OperationContract]
        Task<int> CreateInvoice(Invoice invoice);
    }

    public class SoapService : ISoapService
    {
        private readonly string _connectionString = Environment.GetEnvironmentVariable("connection");
        private readonly IHttpContextAccessor _httpContextAccessor;
        public SoapService(IHttpContextAccessor httpContextAccessor)
        {
            _httpContextAccessor = httpContextAccessor;
        }
        public async Task<int> CreateInvoice(Invoice invoice)
        {
            if (invoice == null)
            {
                throw new ArgumentNullException(nameof(invoice), "Invoice cannot be null.");
            }

            var context = _httpContextAccessor.HttpContext;
            string token = context?.Request.Cookies["token"];

            if (string.IsNullOrEmpty(token))
            {
                Console.WriteLine("Auth Token not found.");
                return -1;
            }

            AppContext.SetSwitch("Switch.Microsoft.Data.SqlClient.DisableSqlServerPerformanceCounters", true);

            string serverAddress = Environment.GetEnvironmentVariable("auth");

            using (var channel = GrpcChannel.ForAddress(serverAddress))
            {
                var client = new AuthService.AuthService.AuthServiceClient(channel);

                var request = new ValidateTokenRequest
                {
                    Token = token,
                    RequiredRole = "admin" // TODO: manage roles
                };
                Console.WriteLine(token);

                try
                {
                    var response = await client.ValidateTokenAsync(request);

                    if (!response.Valid)
                    {
                        Console.WriteLine($"Token validation failed: {response.Message}");
                        return -1; 
                    }

                    Console.WriteLine($"Token valid: {response.Valid}");
                    Console.WriteLine($"User ID: {response.UserId}");
                }
                catch (Grpc.Core.RpcException ex)
                {
                    Console.WriteLine($"Error calling gRPC: {ex}");
                    return -1;
                }
            }

            try
            {
                using var connection = new SqlConnection(_connectionString);
                using var command = new SqlCommand("CreateInvoice", connection)
                {
                    CommandType = CommandType.StoredProcedure
                };

                command.Parameters.AddWithValue("@billing_date", invoice.BillingDate);
                command.Parameters.AddWithValue("@payment_method", invoice.PaymentMethod);
                command.Parameters.AddWithValue("@amount", invoice.Amount);
                command.Parameters.AddWithValue("@shipping_fee", invoice.ShippingFee);
                command.Parameters.AddWithValue("@client_id", invoice.ClientId);
                command.Parameters.AddWithValue("@tax", invoice.Tax);

                connection.Open();
                var result = await command.ExecuteScalarAsync();
                return Convert.ToInt32(result);
            }
            catch (SqlException ex)
            {
                Console.WriteLine($"SQL Error: {ex.Message}");
                return -1;
            }
        }
    }

}